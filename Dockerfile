FROM golang:alpine3.19

EXPOSE 8080

ADD . /app
WORKDIR /app

RUN go build -o app ./cmd

CMD ["./app"]