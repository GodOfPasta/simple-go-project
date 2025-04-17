FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY service/go.mod service/go.sum ./
RUN go mod download

COPY ./service .

RUN CGO_ENABLED=1 GOOS=linux go build -o simple-go-project main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/simple-go-project .
EXPOSE 8080
CMD ["./simple-go-project"]