build:
	@go build -o bin/nse-india-go

run: build
	@./bin/nse-india-go

test:
	@go test -v ./...