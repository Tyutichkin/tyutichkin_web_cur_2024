FROM golang:1.20-alpine AS builder

RUN apk --no-cache add bash git make gcc gettext musl-dev
WORKDIR /app

COPY ./backend .
RUN go build -o app ./cmd/main/main.go
ENTRYPOINT [ "./app" ]

EXPOSE 8083

FROM alpine
COPY --from=builder /app /

CMD ["/app"]