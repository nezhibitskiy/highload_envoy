##
## Build
##
#FROM golang:1.19 AS build
#
#WORKDIR /build
#
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
#
#COPY . .
#
#RUN CGO_ENABLED=0 go build ./main.go

##
## Deploy
##

FROM ubuntu:20.04

WORKDIR /app

COPY ./main .
COPY ./.env .

RUN chmod +x ./main

EXPOSE 5000/tcp

ENTRYPOINT ["./main"]
