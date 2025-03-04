FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/bin/app cmd/main.go

FROM gcr.io/distroless/static-debian10

WORKDIR /root/

COPY --from=builder /app/bin/app .
COPY .env .

CMD ["./app", "-config-path=.env"]
