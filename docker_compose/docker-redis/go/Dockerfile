FROM golang:latest AS build

COPY ./src/main.go /usr/go/src/
RUN go get github.com/gomodule/redigo/redis
WORKDIR /usr/go/src
CMD ["go", "run", "main.go"]