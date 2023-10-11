# build stage
FROM golang:1.18-alpine AS build-go
RUN apk --no-cache add git gcc musl-dev
RUN apk --update add ca-certificates

ENV CODE_PATH /go/src/github.com/velopaymentsapi/velo-go

RUN mkdir -p ${CODE_PATH}
COPY . ${CODE_PATH}
WORKDIR ${CODE_PATH}
RUN go mod tidy