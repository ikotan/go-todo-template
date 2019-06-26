FROM golang:1.12-alpine

WORKDIR /go/src/github.com/ikotan/go-todo-template/

ENV GO111MODULE=on

RUN apk --no-cache --update upgrade \
    && apk add --no-cache git alpine-sdk \
    && go get -u github.com/codegangsta/gin

# RUN apk --no-cache --update upgrade \
#     && apk add --no-cache git alpine-sdk \
#     && go get github.com/pilu/fresh

CMD gin -i run main.go