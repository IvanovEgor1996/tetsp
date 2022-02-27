FROM golang:1.15.5-alpine3.12
LABEL maintainer="Ivanov"
WORKDIR /app
COPY . /app
RUN apk --no-cache update \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates \
    && go mod download \
    && apk add bash

