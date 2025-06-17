FROM ubuntu:22.04

ENV TZ="Europe/Warsaw"
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y software-properties-common && \
    apt-get update

RUN apt install -y golang-go

COPY main.go ./

RUN go mod init server

RUN go get github.com/rs/cors

EXPOSE 8080

CMD go run main.go
