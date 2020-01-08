FROM alpine:latest

RUN mkdir /serpent_aes
WORKDIR /serpent_aes

RUN apk update
RUN apk add musl-dev
RUN apk add go
RUN apk add git

ENV GOPATH /root/go
ENV PATH $PATH:$GOPATH/bin

RUN go get github.com/urfave/cli
RUN go get github.com/kisielk/godepgraph
