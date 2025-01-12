FROM golang:alpine3.21 as builder

WORKDIR /golangfromzero

COPY . .

RUN apk update && apk upgrade

ENTRYPOINT [ "/bin/sh" ]
CMD ["-c", "while true; do echo 'Goland from zero...'; sleep 300; done"]
