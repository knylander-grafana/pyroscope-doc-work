FROM golang:1.22.11

WORKDIR /go/src/app

COPY main.go go.mod go.sum ./

RUN go get -d ./
RUN go build -o main .

RUN adduser --disabled-password --gecos --quiet pyroscope
USER pyroscope

CMD ["./main"]
