FROM golang:alpine

WORKDIR /escapade
COPY go.mod .

RUN apk add --update git
RUN apk add --update bash && rm -rf /var/cache/apk/*
RUN go mod download

COPY . .
