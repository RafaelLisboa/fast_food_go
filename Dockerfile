FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/server

RUN go build -o /app/myapp

WORKDIR /app

EXPOSE 8080

CMD ["./myapp"]
