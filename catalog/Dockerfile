FROM golang

RUN mkdir -p /go/src/github.com/jmvargas/go-kit-example

ADD . /go/src/github.com/jmvargas/go-kit-example

ENTRYPOINT  go run /go/src/github.com/jmvargas/go-kit-example/catalog/cmd/main.go
