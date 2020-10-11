FROM golang:latest AS build

COPY ./src/main.go /usr/go/src/
WORKDIR /usr/go/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM alpine:latest as run

COPY --from=build /usr/go/src/main /usr/go/main
WORKDIR /usr/go
CMD ["./main"]