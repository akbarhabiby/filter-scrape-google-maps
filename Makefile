NAME=filter_scrape

build:
	@echo "Building..."
	go mod tidy
	CGO_ENABLED=0 go build -o bin/${NAME}
	@echo "Build success!"

build-windows-amd64:
	@echo "Building..."
	go mod tidy
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/${NAME}
	@echo "Build success!"

.PHONY: build build-windows-amd64
