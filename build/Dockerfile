FROM golang:1.16 AS builder

WORKDIR /talent-acquisition-management

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -o cli.bin cmd/main.go

FROM alpine:latest

RUN apk update && apk add bash

COPY --from=builder /talent-acquisition-management/db/migrations ./migrations
COPY --from=builder /talent-acquisition-management/build/wait-for-it.sh ./wait-for-it.sh

# Binary file
COPY --from=builder /talent-acquisition-management/cli.bin ./
