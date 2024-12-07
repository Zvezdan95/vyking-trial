FROM golang:1.23.4-alpine

WORKDIR /usr/src/app

# Install air
RUN go install github.com/air-verse/air@latest

# Copy the rest of the project files
COPY . .
RUN go mod tidy