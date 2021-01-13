FROM golang:alpine as builder

LABEL maintainer="Bryan Builes <brayam.builes@gmail.com>"

WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# COPY --from=builder /app/.env .
# Expose port 8080 to the outside world
EXPOSE 3001
# Command to run the executable
CMD ["./main"]
