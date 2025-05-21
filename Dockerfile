FROM golang:1.24.0 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM alpine:latest  
WORKDIR /root/

COPY --from=builder /docker-gs-ping .
COPY --from=builder /app/.env .env

CMD ["./docker-gs-ping"]
