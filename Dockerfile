FROM golang:1.17.9-alpine

WORKDIR /build_dir

ENV GO111MODULE=on

COPY . .

RUN go build -o mail main.go

ENTRYPOINT ["/build_dir/mail"]
