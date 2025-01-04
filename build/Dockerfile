FROM golang:alpine3.21 as builder

WORKDIR /golangdozero

COPY . .

RUN apk update && apk upgrade

ENTRYPOINT [ "/bin/sh" ]
CMD ["-c", "while true; do echo 'Aprenda Goland do zero...'; sleep 300; done"]
