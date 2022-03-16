FROM golang:1.17 AS builder
LABEL maintainer="eco@economicus.kr"
LABEL version="1.0.0"
LABEL description="Economicus main server"

RUN apt-get update && apt-get install -y wget && apt install -y protobuf-compiler

ENV GOPATH /go
ENV PATH $PATH:/go/bin:$GOPATH/bin
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN wget https://raw.githubusercontent.com/economicus/backend-proto/main/go/quant.pb.go -O internal/core/pb/quant.pb.go && \
		wget https://raw.githubusercontent.com/economicus/backend-proto/main/go/quant_grpc.pb.go -O internal/core/pb/quant_grpc.pb.go


RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/economicus/*.go

FROM alpine:latest AS production

COPY --from=builder /app .
