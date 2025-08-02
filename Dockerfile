FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM alpine:3.21.3 AS final

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main /app
COPY --from=builder /app/.env /app

EXPOSE 8080

CMD ["./main"]
