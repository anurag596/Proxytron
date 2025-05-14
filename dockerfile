# Use the official Golang image to build the binary
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o proxytron .

# Final image
FROM alpine:latest

COPY --from=builder /app/proxytron /proxytron

EXPOSE 8080

ENTRYPOINT ["/proxytron"]
