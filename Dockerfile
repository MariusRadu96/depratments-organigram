FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o department-hierarchy ./cmd/departments

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/department-hierarchy /app/department-hierarchy

EXPOSE 8080

CMD ["/app/department-hierarchy"]