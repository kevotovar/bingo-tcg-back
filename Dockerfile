FROM golang:alpine

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/pilu/fresh

ENTRYPOINT fresh
