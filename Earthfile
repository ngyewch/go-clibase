VERSION 0.7
FROM golang:1.19.7-alpine
WORKDIR /workspace

build:
    # Download deps before copying code.
    COPY go.mod go.sum .
    RUN go mod download
    # Copy and build code.
    COPY . .
    RUN go build ./...
