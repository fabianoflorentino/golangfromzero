FROM golang:alpine3.21 as builder

WORKDIR /golangdozero

COPY . .

RUN go build -o /go/bin/golangdozero cmd/golangdozero/main.go

FROM alpine:3.21 as golangdozero

COPY --from=builder /go/bin/golangdozero /go/bin/golangdozero

RUN apk update && apk upgrade --no-cache

ENTRYPOINT [ "/bin/sh" ]
CMD ["-c", "while true; do echo 'Aprenda Goland do zero...'; sleep 300; done"]
