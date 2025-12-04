FROM golang:1.25.5-alpine AS builder

ARG SERVICE_NAME
ARG VERSION
ARG COMMIT
ARG BUILD_DATE

RUN apk add --no-cache git

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -tags netgo \
    -ldflags "\
        -extldflags '-static' \
        -X 'main.date=${BUILD_DATE}' \
        -X 'main.name=${SERVICE_NAME}' \
        -X 'main.version=${VERSION}' \
        -X 'main.commit=${COMMIT}'" \
    -o ./build/app ./cmd/.

FROM alpine:3.23.0

ARG SERVICE_NAME

RUN apk --no-cache add ca-certificates wget
WORKDIR /app

COPY --from=builder /app/build/app .
COPY --from=builder /app/cmd/config.yaml ./cmd/config.yaml

EXPOSE 10000
ENTRYPOINT ["./app"]
