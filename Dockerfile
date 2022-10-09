FROM golang:1.17.9-alpine

COPY . .

ENTRYPOINT ["go","run","main.go"]
