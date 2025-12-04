.EXPORT_ALL_VARIABLES:
.PHONY: help up down destroy check-env

COMPOSE_FILE := docker-compose.yml
PROJECT_NAME := guru-apps-test-services
NOTIFICATIONS_ENV := notifications-service/.env
PRODUCTS_ENV := products-service/.env
NOTIFICATIONS_EXAMPLE := notifications-service/example_docker.env
PRODUCTS_EXAMPLE := products-service/example_docker.env

help:
	@echo "\033[1;33m🐳 Docker Compose commands:\033[0m"
	@echo "\033[1;34m  make up              \t Start all services (auto-creates .env if needed)\033[0m"
	@echo "\033[1;34m  make down            \t Stop and remove containers\033[0m"
	@echo "\033[1;34m  make destroy         \t Destroy everything (containers, images, volumes, networks)\033[0m"
	@echo ""

check-env:
	@if [ ! -f $(NOTIFICATIONS_ENV) ]; then \
		echo "\033[1;33m⚠️  File $(NOTIFICATIONS_ENV) not found.\033[0m"; \
		printf "\033[1;36m   Create it from $(NOTIFICATIONS_EXAMPLE)? [y/N]: \033[0m"; \
		read answer; \
		if [ "$$answer" = "y" ] || [ "$$answer" = "Y" ]; then \
			cp $(NOTIFICATIONS_EXAMPLE) $(NOTIFICATIONS_ENV); \
			echo "\033[1;32m✓ Created $(NOTIFICATIONS_ENV)\033[0m"; \
		else \
			echo "\033[1;31m✗ Aborted. Please create $(NOTIFICATIONS_ENV) manually.\033[0m"; \
			exit 1; \
		fi; \
	fi
	@if [ ! -f $(PRODUCTS_ENV) ]; then \
		echo "\033[1;33m⚠️  File $(PRODUCTS_ENV) not found.\033[0m"; \
		printf "\033[1;36m   Create it from $(PRODUCTS_EXAMPLE)? [y/N]: \033[0m"; \
		read answer; \
		if [ "$$answer" = "y" ] || [ "$$answer" = "Y" ]; then \
			cp $(PRODUCTS_EXAMPLE) $(PRODUCTS_ENV); \
			echo "\033[1;32m✓ Created $(PRODUCTS_ENV)\033[0m"; \
		else \
			echo "\033[1;31m✗ Aborted. Please create $(PRODUCTS_ENV) manually.\033[0m"; \
			exit 1; \
		fi; \
	fi

up: check-env
	@echo "\033[1;32m🚀 Starting all services...\033[0m"
	docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) up --build

down:
	@echo "\033[1;33m🛑 Stopping all services...\033[0m"
	docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down

destroy:
	@echo "\033[1;31m🗑️  Destroying all containers, images, volumes, and networks...\033[0m"
	-docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down -v --rmi all --remove-orphans
	-docker volume prune -f
	-docker network prune -f
	@echo "\033[1;32m✓ All artifacts destroyed\033[0m"
