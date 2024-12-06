# Makefile for SlotMachine Web Service Project

# Default to run the server
run:
	go run main.go

# Build the executable
build:
	go build -o web-service main.go

# Test the project
test:
	go test ./...

# Clean up built files
clean:
	rm -f web-service
