FROM golang:1.15.7-alpine3.13

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/brkelkar/awacs_smart_api_service

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

ENV GIN_MODE=release
ENV SERVER_PORT=8080

# Install the package
RUN go build -tags=jsoniter .
# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./awacs_smart_api_service"]