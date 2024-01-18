FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o orderapi order-api/cmd/main.go

CMD ["./orderapi"]