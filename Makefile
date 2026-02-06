.PHONY: build run-ingestion clean

# Build the Ingestion Service binary
build: 
	@echo "Building Ingestion Service..."
	@go build -o bin/ingestion-api ./cmd/ingestion-api/main.go

# Run the Ingestion Service directly
run-ingestion: 
	@echo "Running Ingestion Service..."
	@go run ./cmd/ingestion-api/main.go

#Clean up binaries
	@echo "Cleaning..."
	@rm -rf bin

