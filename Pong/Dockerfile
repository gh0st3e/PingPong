FROM golang:alpine AS builder

COPY Pong/. /src
WORKDIR /src
RUN go build -o main.bin cmd/main.go

FROM alpine AS runner

WORKDIR /app
COPY --from=builder /src/main.bin /app/

ENTRYPOINT ./main.bin

CMD ["/app"]