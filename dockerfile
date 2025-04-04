FROM golang:1.22-alpine AS builder

RUN apk --no-cache add \
    bash=5.2.37-r0 \
    git=2.47.2-r0 \
    make=4.4.1-r2 \
    gcc=14.2.0-r4 \
    gettext=0.22.5-r0 \
    musl-dev=1.2.5-r9 \
    curl=8.12.1-r1 \
    postgresql17-client=17.4-r0

WORKDIR /app

COPY ./backend .
RUN go build -o app ./cmd/main/main.go

# Установка goose
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.20.0 && \
    ln -s "$(go env GOPATH)/bin/goose" /usr/local/bin/goose

COPY backend/migrate_and_run.sh /app/migrate_and_run.sh
RUN chmod +x /app/migrate_and_run.sh

ENTRYPOINT ["/app/migrate_and_run.sh"]
EXPOSE 8083

FROM alpine:3.20

RUN apk --no-cache add \
    git=~2.45.3-r0 \
    bash=~5.2.26-r0 \
    curl=~8.12.1-r0 \
    postgresql15-client=~15.11-r0

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

COPY --from=builder /app /
CMD ["/app"]