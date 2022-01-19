FROM golang:1.17-alpine3.15 AS builder
WORKDIR /root/src
COPY . /root/src/
RUN go build -o ./build/tetrominos .

FROM alpine:3.15.0
WORKDIR /root/
COPY --from=builder /root/src/build/tetrominos ./