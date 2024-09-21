# Variables
APP_NAME=my-go-app
DOCKER_IMAGE=my-go-app-image

# Build the Go application
build:
	@go build -o main .

# Build Docker Image
docker-build:
	@docker build -t $(DOCKER_IMAGE) .

# Run Docker Container
docker-run:
	@docker run --rm -p 8080:8080 $(DOCKER_IMAGE)

# Clean up
clean:
	@rm -f main