FROM golang:1.12.6-stretch

ENV GO111MODULE=on

WORKDIR /go/src/app

RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get github.com/Shelnutt2/db2struct/db2struct
