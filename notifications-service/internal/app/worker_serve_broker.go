package app

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

// Event types for notifications services.
const (
	eventTypeCreateProduct = "create_product"
	eventTypeDeleteProduct = "delete_product"
)

// brokerRoutes registers broker routes.
func (a *App) brokerHandlers() map[string]func(ctx context.Context, body []byte) error {
	return map[string]func(ctx context.Context, body []byte) error{
		eventTypeCreateProduct: a.sqsConsumerHandler.CreateNotification,
		eventTypeDeleteProduct: a.sqsConsumerHandler.DeleteNotification,
	}
}

// serveBroker listen for registered subjects.
func serveBroker(ctx context.Context, app *App) {
	queueURL := app.cfg.Delivery.Broker.URL
	handlers := app.brokerHandlers()

	app.logger.Info("starting SQS consumer")

	wg := &sync.WaitGroup{}

	go func() {
		for {
			select {
			case <-ctx.Done():
				app.logger.Info("stopping SQS consumer, waiting for messages to finish processing")
				wg.Wait()
				app.logger.Info("SQS consumer stopped gracefully")
				return
			default:
				resp, err := app.sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
					QueueUrl:            &queueURL,
					MaxNumberOfMessages: app.cfg.Delivery.Broker.MaxNumberOfMessages,
					WaitTimeSeconds:     app.cfg.Delivery.Broker.WaitTimeSeconds,
				})
				if err != nil {
					if ctx.Err() != nil {
						return
					}
					app.logger.Error("failed to receive messages", zap.Error(err))
					time.Sleep(app.cfg.Delivery.Broker.RetryDelay)
					continue
				}

				if resp == nil || len(resp.Messages) == 0 {
					continue
				}

				for _, m := range resp.Messages {
					select {
					case <-ctx.Done():
						return
					default:
					}

					wg.Add(1)
					go func(msg types.Message) {
						defer wg.Done()

						handlerCtx, handlerCancel := context.WithTimeout(ctx, app.cfg.Delivery.Broker.HandlerTimeout)
						defer handlerCancel()

						delCtx, delCancel := context.WithTimeout(context.Background(),
							app.cfg.Delivery.Broker.DeleteTimeout)
						defer delCancel()

						if msg.Body == nil || *msg.Body == "" {
							app.logger.Warn("received message with empty body")
							return
						}

						var msgBody struct {
							EventType string `json:"event_type"`
						}
						if err = json.Unmarshal([]byte(*msg.Body), &msgBody); err != nil {
							app.logger.Error("failed to parse message", zap.Error(err),
								zap.String("message_body", *msg.Body))
							if _, delErr := app.sqsClient.DeleteMessage(delCtx, &sqs.DeleteMessageInput{
								QueueUrl:      &queueURL,
								ReceiptHandle: msg.ReceiptHandle,
							}); delErr != nil {
								app.logger.Error("failed to delete invalid message", zap.Error(delErr))
							}
							return
						}

						handler, ok := handlers[msgBody.EventType]
						if !ok {
							app.logger.Warn("no handler for event type",
								zap.String("event_type", msgBody.EventType))
							if _, delErr := app.sqsClient.DeleteMessage(delCtx, &sqs.DeleteMessageInput{
								QueueUrl:      &queueURL,
								ReceiptHandle: msg.ReceiptHandle,
							}); delErr != nil {
								app.logger.Error("failed to delete unhandled message", zap.Error(delErr))
							}
							return
						}

						if err = handler(handlerCtx, []byte(*msg.Body)); err != nil {
							app.logger.Error("handler failed", zap.Error(err),
								zap.String("event_type", msgBody.EventType),
								zap.String("message_body", *msg.Body))
							return
						}

						if _, err = app.sqsClient.DeleteMessage(delCtx, &sqs.DeleteMessageInput{
							QueueUrl:      &queueURL,
							ReceiptHandle: msg.ReceiptHandle,
						}); err != nil {
							app.logger.Error("failed to delete msg after successful processing", zap.Error(err))
						}
					}(m)
				}
			}
		}
	}()
}
