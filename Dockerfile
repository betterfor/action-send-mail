FROM golang:1.17.9

COPY . .

ENTRYPOINT ["go","run","main.go"]
