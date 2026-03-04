FROM golang:alpine3.23.3 AS base

WORKDIR /app

COPY . .

RUN go mod download \
  && GOFLAGS="-trimpath" CGO_DISABLED=1 GOARCH=amd64 \
  go build -ldflags="-s -w" -o /usr/local/bin/app /app/cmd/app


FROM base AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest

EXPOSE 8080

ENTRYPOINT [ "/go/bin/air" ]
CMD [ "-c", "/app/.air.toml" ]


FROM gcr.io/distroless/static:nonroot AS production

COPY --from=base /usr/local/bin/app /usr/local/bin/app

USER nonroot:nonroot

ENTRYPOINT [ "/usr/local/bin/app" ]
