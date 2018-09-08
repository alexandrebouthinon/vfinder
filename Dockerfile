FROM golang:latest

RUN go get github.com/alexandrebouthinon/vfinder && go install github.com/alexandrebouthinon/vfinder

