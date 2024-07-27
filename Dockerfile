FROM golang:alpine3.19

EXPOSE 8080

ADD . /app
WORKDIR /app

RUN apk add --no-cache git
RUN go build -ldflags="-X main.Commit=$(git rev-parse --short HEAD)" -o app ./cmd

CMD ["./app"]