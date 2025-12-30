FROM golang:alpine3.22 AS base

WORKDIR /golangfromzero

COPY . .

RUN apk update -y --no-cache \
  && apk upgrade -y --no-cache \
  && go mod download \
  && GOFLAGS="-trimpath" CGO_DISABLED=1 GOARCH=amd64 go build -ldflags="-s -w" -o /usr/local/bin/golangfromzero /golangfromzero/cmd/golangfromzero

FROM base AS development

RUN go install github.com/air-verse/air@latest

COPY --from=base /golangfromzero/config/.env.example /root/.env

EXPOSE 6000

ENTRYPOINT [ "/go/bin/air" ]
CMD [ "-c", "/golangfromzero/.air.toml" ]


FROM gcr.io/distroless/static:nonroot AS production

COPY --from=base /usr/local/bin/golangfromzero /usr/local/bin/golangfromzero

USER nonroot:nonroot

ENTRYPOINT [ "/usr/local/bin/golangfromzero" ]