ARG BUILD_GO_VERSION=1.24
ARG BUILD_OS_RELEASE=alpine
ARG RUN_OS_VERSION=3.21

FROM golang:${BUILD_GO_VERSION}-${BUILD_OS_RELEASE} AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/app ./cmd

# Runner
FROM alpine:${RUN_OS_VERSION}

COPY --from=builder /bin/app /app
COPY config.yaml ./

EXPOSE 8080

ENTRYPOINT ["/app"]