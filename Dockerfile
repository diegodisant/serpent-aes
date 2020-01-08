FROM alpine:latest

RUN mkdir /serpent_aes
VOLUME /serpent_aes
WORKDIR /serpent_aes
ADD . /serpent_aes

RUN apk update
RUN apk add musl-dev
RUN apk add go
RUN apk add git

ENV GOPATH /root/go
ENV PATH $PATH:$GOPATH/bin

RUN go get github.com/urfave/cli
RUN go get github.com/kisielk/godepgraph
RUN sh setup-docker.sh
RUN cd test && go test -test.v
