FROM golang:1.10

MAINTAINER jackliu97@gmail.com

WORKDIR /tmp

COPY . /go/src/chatter

WORKDIR /go/src/chatter

RUN go get && go build .

ENTRYPOINT ./chatter

EXPOSE 8080 8080

