FROM golang:1.10

MAINTAINER jackliu97@gmail.com

WORKDIR /app

COPY . /go/src/github.com/jackliu97/chatter

WORKDIR /go/src/github.com/jackliu97/chatter

RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/spf13/viper

RUN go install

ENTRYPOINT /go/bin/chatter

EXPOSE 8080
