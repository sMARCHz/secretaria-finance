FROM golang:1.19-alpine

ENV CGO_ENABLED=0

WORKDIR /app

RUN go install github.com/cespare/reflex@latest