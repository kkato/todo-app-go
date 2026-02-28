FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.ini /app/config.ini
COPY --from=builder /app/app/views /app/app/views

EXPOSE 8080

CMD ["/app/main"]
