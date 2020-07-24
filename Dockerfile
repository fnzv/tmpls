FROM golang:1.10

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./trackmepls.go

CMD ["./trackmepls"]

