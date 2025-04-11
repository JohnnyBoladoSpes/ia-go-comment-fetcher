.PHONY: up down logs build dev test clean status

up:
	@echo "🚀 Starting all services..."
	docker compose down && docker compose up --build -d

down:
	@echo "🛑 Stopping all services..."
	docker compose down

logs:
	@echo "📜 Showing logs for Go service..."
	docker logs -f ia-go-comment-fetcher

build:
	@echo "🔨 Building Go binary..."
	go build -o comment-fetcher .

dev: deps
	@echo "🚀 Starting Mongo & Redis via Docker..."
	docker compose up -d mongo redis
	@echo "🔁 Running Go service with hot reload (air)..."
	air

test:
	@echo "🧪 Running tests..."
	go test ./...

clean:
	@echo "🧹 Cleaning up Docker and cache files..."
	docker system prune -af
	find . -name "__pycache__" -exec rm -rf {} +
	rm -f comment-fetcher

status:
	@echo "Checking status of Go service (port 8080)..."
	@if nc -zv 127.0.0.1 8080 > /dev/null 2>&1; then \
		echo "✅ Go service is running"; \
	else \
		echo "❌ Go service is not running"; \
	fi
	@echo "Checking status of MongoDB..."
	@if nc -zv 127.0.0.1 27018 > /dev/null 2>&1; then \
		echo "✅ MongoDB is running"; \
	else \
		echo "❌ MongoDB is not running"; \
	fi
	@echo "Checking status of Redis..."
	@if nc -zv 127.0.0.1 6380 > /dev/null 2>&1; then \
		echo "✅ Redis is running"; \
	else \
		echo "❌ Redis is not running"; \
	fi

deps:
	@echo "📦 Installing Go dependencies..."
	go mod tidy
	go mod download

deps-clean:
	@echo "🧼 Cleaning and reinstalling Go dependencies..."
	go clean -modcache
	go mod tidy
	go mod download