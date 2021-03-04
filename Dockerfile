FROM golang:1.15.7-alpine3.13

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

ENV GIN_MODE=release
ENV SERVER_PORT=8080

# Install the package
RUN go build -tags=jsoniter .
# This container exposes port 8080 to the outside world

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /root/awacs_smart_api_service . 
EXPOSE 8080
# Run the executable
CMD ["./awacs_smart_api_service"]
