run:
	@go run cmd/main/main.go

build:
	@go build -o bin/fetch-backend-service cmd/main/main.go
