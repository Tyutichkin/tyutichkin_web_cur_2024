FROM golang:1.23-alpine AS builder

RUN apk --no-cache add bash git make gcc gettext musl-dev curl postgresql-client

WORKDIR /app

COPY ./backend .
RUN go build -o app ./cmd/main/main.go
# Установка goose
RUN go get -u github.com/pressly/goose && \
    ln -s $(go env GOPATH)/bin/goose /usr/local/bin/goose


COPY backend/migrate_and_run.sh /app/migrate_and_run.sh
RUN chmod +x /app/migrate_and_run.sh

ENTRYPOINT ["/app/migrate_and_run.sh"]
EXPOSE 8083

FROM alpine

RUN apk --no-cache add git bash curl postgresql-client
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

COPY --from=builder /app /
CMD ["/app"]