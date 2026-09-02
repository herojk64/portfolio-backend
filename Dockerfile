FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/api ./cmd/api
RUN go build -o /app/seed ./cmd/seed

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/api /app/api
COPY --from=builder /app/seed /app/seed
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

COPY config/ ./config/
COPY sql/migrations/ ./sql/migrations/
COPY internal/seed/data/ ./internal/seed/data/
COPY entrypoint.sh .

RUN chmod +x /app/entrypoint.sh

EXPOSE 8000

ENTRYPOINT ["/app/entrypoint.sh"]
