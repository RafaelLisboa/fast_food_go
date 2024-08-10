FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY ./build/ ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/server
RUN go build -o /app/myapp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp /app/myapp

EXPOSE 8080

CMD ["./myapp"]

