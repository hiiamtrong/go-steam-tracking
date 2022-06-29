FROM golang:alpine AS builder
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
WORKDIR /src/local/app

ENV TZ=Asia/Ho_Chi_Minh

COPY  go.mod .
COPY  go.sum .
RUN go mod download
COPY . .

RUN go build -o bin/server cmd/server/main.go
RUN go build -o bin/crawler cmd/crawler/main.go

FROM alpine:latest

WORKDIR /src/local/app
RUN apk add -U tzdata
ENV TZ=Asia/Ho_Chi_Minh


COPY --from=builder /src/local/app/bin .
COPY --from=builder /src/local/app/.env .
EXPOSE 3001
