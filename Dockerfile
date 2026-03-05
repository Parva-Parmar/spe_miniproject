FROM golang:1.22-alpine

 WORKDIR /app

 COPY go.mod ./

COPY . .

RUN go build -o calc-app main.go

CMD ["./calc-app"]