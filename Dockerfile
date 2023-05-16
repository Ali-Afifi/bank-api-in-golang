FROM golang:1.17-alpine3.16 AS builder

WORKDIR /app

COPY . .

# ENV PATH=/usr/local/go/bin:$PATH

RUN go mod download && go build -o ./bin/bankserver

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/bin /app/bin

# ENV PORT=8080

# ENV REQUEST_ORIGIN=http://localhost:5000

EXPOSE 8080

CMD ["./bin/bankserver"]
