FROM golang:1.23.2-bookworm

WORKDIR /app

COPY api /app/api

COPY go.mod go.sum main.go /app/

RUN go mod download && go build -o backend

# Run the migrations
RUN ./backend up

# Expose the port
EXPOSE 3001

CMD [ "./backend" ]