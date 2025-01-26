FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.work go.work.sum ./
COPY server/go.mod server/go.sum ./server/
COPY api/go.mod api/go.sum ./api/
COPY auth/go.mod auth/go.sum ./auth/
COPY database/go.mod database/go.sum ./database/
COPY shared/go.mod shared/go.sum ./shared/
RUN go work sync && go mod download all
COPY . .

WORKDIR /app/server
RUN go build -o main .

FROM alpine:latest AS runtime
RUN apk add --no-cache postgresql-client
WORKDIR /app
COPY --from=builder /app/server/main .
COPY --from=builder /app/server/database/migrations ./migrations
EXPOSE ${PORT}
CMD ["./main"]