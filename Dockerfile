FROM nexus.stripchat.tech/golang:1.19 as builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/go-service ./cmd/app/main.go

FROM scratch
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/go-service /app/go-service
COPY --from=builder /app/api/. /app/api/.
COPY --from=builder /app/config.yaml /app/config.yaml

ENTRYPOINT ["/app/go-service"]
