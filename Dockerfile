FROM golang:1.18-alpine

WORKDIR /app
COPY . .

CMD ["go run main.go"]