FROM chatter-base

COPY . /go/src/github.com/jackliu97/chatter

WORKDIR /go/src/github.com/jackliu97/chatter

RUN go install

ENTRYPOINT /go/bin/chatter

EXPOSE 8080
