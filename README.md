# Test Services

Microservices application built with Go, featuring products and notifications services with PostgreSQL, LocalStack (SQS), and Prometheus monitoring.

## 🎯 Overview

This project consists of two microservices:

- **products-service** - Manages product data with PostgreSQL database
- **notifications-service** - Handles notifications via SQS queue

Both services are containerized using Docker and orchestrated with Docker Compose.

## 🏗️ Architecture

```
┌─────────────────┐     ┌──────────────────┐
│  Products       │────▶│   PostgreSQL     │
│  Service        │     │   (Port 5432)    │
│  (Port 10000)   │     └──────────────────┘
└────────┬────────┘
         │
         │ SQS
         ▼
┌─────────────────┐     ┌──────────────────┐
│ Notifications   │────▶│   LocalStack     │
│ Service         │     │   (Port 4566)    │
│ (Port 10001)    │     └──────────────────┘
└─────────────────┘
         │
         │ Metrics
         ▼
┌─────────────────┐
│   Prometheus    │
│   (Port 9090)   │
└─────────────────┘
```

## 🚀 Quick Start

### Start All Services

The Makefile will automatically check for required `.env` files and prompt you to create them from examples if they don't exist.

```bash
make up
```

This command will:
- Check for `.env` files in both service directories
- Prompt you to create them from `example_docker.env` if missing
- Build and start all services (PostgreSQL, Prometheus, etc.)


## 🛠️ Available commands

### Using Makefile

```bash
# Show help
make help
```
```bash
# Start all services (auto-creates .env files if needed)
make up
```
```bash
# Stop all services
make down
```
```bash
# Destroy everything (containers, images, volumes, networks)
make destroy
```

## 📡 API Endpoints

### Products Service

Base URL: `http://localhost:10000/products-api/v1`

| Method     | Endpoint        | Description                            |
|------------|-----------------|----------------------------------------|
| **GET**    | `/health`       | Health check                           |
| **GET**    | `/products`     | Get all products with limit and offset |
| **POST**   | `/products`     | Create a new product                   |
| **DELETE** | `/products/:id` | Delete a product                       |
| **GET**    | `/metrics`      | Prometheus metrics                     |

### Notifications Service

Base URL: `http://localhost:10001/notifications-api/v1`

| Method | Endpoint  | Description  |
|--------|-----------|--------------|
| GET    | `/health` | Health check |

## ⚙️ Configuration

### Environment Variables

Each service requires a `.env` file. Example files are provided:

- `notifications-service/example_docker.env`
- `products-service/example_docker.env`

The Makefile will automatically prompt you to create `.env` files from examples if they don't exist.

## 🧪 Testing

### Unit Tests

Integration tests for the products repository are available in `products-service/tests/integration/products_repository_test.go`:

- Creation
- Retrieval 
- Deletion

#### Running the Tests

To run the tests, ensure PostgreSQL is running and execute:

```bash
cd products-service
go test ./tests/integration/... -v
```

### Testing the API

#### 1. Using HTTP Files (JetBrains IDE)

Run the HTTP files:
- `products-service/fixtures/http/`

#### 2. Using Postman

Import the Postman collection from:
- `products-service/fixtures/postman/postman_collection.json`

#### 3. Using cURL

```bash
# Health check
curl http://localhost:10000/products-api/v1/health | jq

# Get all products
curl http://localhost:10000/products-api/v1/products | jq

# Create a product
curl -X POST http://localhost:10000/products-api/v1/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Product", "vendor": "Test Vendor", "price": 99.99}' | jq
```

## 🔍 Monitoring

### Prometheus

Prometheus is configured to scrape metrics. Access the Prometheus UI at: **`http://localhost:9090`**

#### Available Metrics

The following metrics are exposed by the services:

- `products_service_created_products_cnt` - Total number of products created
- `products_service_deleted_products_cnt` - Total number of products deleted

#### Quick Links

- **Prometheus UI**: [http://localhost:9090](http://localhost:9090)
- **View Example Metrics**: [Products Created & Deleted](http://localhost:9090/query?g0.expr=products_service_created_products_cnt&g0.show_tree=0&g0.tab=graph&g0.range_input=5m&g0.res_type=auto&g0.res_density=high&g0.display_mode=lines&g0.show_exemplars=0&g1.expr=products_service_deleted_products_cnt&g1.show_tree=0&g1.tab=graph&g1.range_input=5m&g1.res_type=auto&g1.res_density=high&g1.display_mode=lines&g1.show_exemplars=0)

#### Example Queries

You can use the Prometheus query language (PromQL) to explore metrics:

```promql
# Total products created
products_service_created_products_cnt

# Total products deleted
products_service_deleted_products_cnt
```
