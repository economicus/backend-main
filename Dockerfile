FROM golang:1.17 AS builder
LABEL maintainer="eco@economicus.kr"
LABEL version="1.0.0"
LABEL description="Economicus main server"

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/economicus/*.go

FROM alpine:latest AS production

COPY --from=builder /app .