FROM golang:1.10.3-alpine3.7

COPY . /go/src/github.com/rajugupta15/nats-streaming-producer
WORKDIR /go/src/github.com/rajugupta15/nats-streaming-producer
RUN apk add git

RUN go get ; CGO_ENABLED=0 GOOS=linux go build -o nats-streaming-producer main.go
FROM alpine:3.8
RUN apk add ca-certificates
COPY --from=0 /go/src/github.com/rajugupta15/nats-streaming-producer/nats-streaming-producer /nats-streaming-producer
CMD [ "/nats-streaming-producer" ]