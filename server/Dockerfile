FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest AS runtime
RUN apk add --no-cache postgresql-client
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/database/migrations ./migrations
EXPOSE ${PORT}
CMD ["./main"]