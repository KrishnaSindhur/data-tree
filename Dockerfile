FROM golang:alpine as builder

ADD . /home

WORKDIR /home

EXPOSE 8080

CMD ["go","run","./cmd/data-tree/main.go", "serve"]
