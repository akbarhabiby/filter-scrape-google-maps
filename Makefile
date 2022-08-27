NAME=filter_scrape

build:
	@echo "Building..."
	go mod tidy
	CGO_ENABLED=0 go build -o bin/${NAME}
	@echo "Build success!"

.PHONY: build
