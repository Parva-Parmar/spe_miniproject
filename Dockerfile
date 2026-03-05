FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o calculator main.go

CMD ["./calculator"]