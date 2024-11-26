FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o backend

FROM builder AS development
CMD ["go", "run", "main.go"]

FROM debian:bookworm-slim AS runner
WORKDIR /app
COPY --from=builder /app/backend ./backend
COPY --from=builder /app/api/images ./api/images
COPY --from=builder /app/api/db/ ./api/db/

# Expose the port
EXPOSE 3001
CMD [ "./backend" ]