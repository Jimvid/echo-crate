# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@templ generate
	@go build -o main cmd/http/main.go

# Run the application
run:
	@go run cmd/http/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload App
watch:
	@if command -v air > /dev/null; then \
	    air;\
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# Watch Tailwind
watch-tailwind:
	echo "Watching Tailwind..."
	./tailwind -i ./assets/css/main.css -o ./assets/css/main.min.css --minify --watch

# Build Tailwind
build-tailwind:
	echo "Building Tailwind..."
	./tailwind -i ./assets/css/main.css -o ./assets/css/main.min.css --minify

# Hot Reload when making changes to templates
templ:
	echo "Watching templ..."
	templ generate -watch -proxy=http://localhost:8080

.PHONY: all build run test clean
