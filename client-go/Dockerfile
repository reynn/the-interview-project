ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} as builder

WORKDIR /app

COPY . .

RUN go build -ldflags="-s" -trimpath -o client cmd/client/client.go

FROM debian

COPY --from=builder /app/client /app/client

CMD ["/app/client"]
