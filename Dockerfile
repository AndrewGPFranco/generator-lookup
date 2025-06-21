FROM golang:1.24.4

WORKDIR /go/src/app

COPY . . 

EXPOSE 8000

RUN go build -o main cmd/api/main.go

CMD ["./main"]