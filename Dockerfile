FROM golang:1.17.9-alpine

WORKDIR /build_dir

ENV GO111MODULE=on

COPY . .

ENTRYPOINT ["go","run","main.go"]
