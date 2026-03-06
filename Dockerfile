FROM golang:alpine3.23 AS base

WORKDIR /app

COPY . .

RUN go mod download \
  && GOFLAGS="-trimpath" CGO_DISABLED=1 GOARCH=amd64 \
  go build -ldflags="-s -w" -o /usr/local/bin/app /app/cmd/app


FROM base AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest \
  && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1 \
  && go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.1

EXPOSE 8080

ENTRYPOINT [ "/go/bin/air" ]
CMD [ "-c", "/app/.air.toml" ]


FROM gcr.io/distroless/static:nonroot AS production

COPY --from=base /usr/local/bin/app /usr/local/bin/app

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT [ "/usr/local/bin/app" ]
