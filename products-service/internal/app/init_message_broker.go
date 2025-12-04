package app

import (
	"context"
	"net/url"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/zap"
)

// initMessageBroker - initialize message broker.
func (a *App) initMessageBroker(ctx context.Context) {
	queueName, err := parseSQSURL(a.cfg.Delivery.Broker.URL)
	if err != nil {
		a.logger.Fatal("cannot parse SQS URL", zap.Error(err))
	}

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(a.cfg.Delivery.Broker.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("test", "test", ""),
		),
	)
	if err != nil {
		a.logger.Fatal("cannot load AWS config", zap.Error(err))
	}

	queueURL, err := url.Parse(a.cfg.Delivery.Broker.URL)
	if err != nil {
		a.logger.Fatal("cannot parse SQS queue URL", zap.Error(err))
	}
	baseEndpointURL := &url.URL{
		Scheme: queueURL.Scheme,
		Host:   queueURL.Host,
	}
	baseEndpoint := aws.String(baseEndpointURL.String())

	a.sqsClient = sqs.NewFromConfig(awsCfg, func(o *sqs.Options) {
		o.BaseEndpoint = baseEndpoint
	})

	if _, err = a.sqsClient.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{QueueName: aws.String(queueName)}); err != nil {
		a.logger.Fatal("cannot get SQS queue URL", zap.Error(err))
	}

	a.logger.Info("SQS LocalStack initialized",
		zap.String("aws_region", a.cfg.Delivery.Broker.Region),
		zap.String("queue_name", queueName))
}

// parseSQSURL - parse SQS URL to get queueName, e.g.:
//   - http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/test-queue
//   - http://localstack:4566/000000000000/test-queue
func parseSQSURL(sqsURL string) (string, error) {
	u, err := url.Parse(sqsURL)
	if err != nil {
		return "", err
	}

	queueName := path.Base(u.Path)

	return queueName, nil
}
