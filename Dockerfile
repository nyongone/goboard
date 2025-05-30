FROM golang:1.24-alpine AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

ENV DOCKERIZE_VERSION=v0.9.3
ENV MIGRATE_VERSION=v4.18.3

RUN apk update --no-cache \
    && apk add --no-cache wget openssl \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && wget -O - https://github.com/golang-migrate/migrate/releases/download/$MIGRATE_VERSION/migrate.linux-amd64.tar.gz | tar xzf - -C /usr/local/bin \
    && apk del wget

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o server main.go

FROM alpine:3.21

COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /app /

RUN chmod +x entrypoint.sh

ENV MIGRATION_DIR=./migrations

ENTRYPOINT ["./entrypoint.sh"]

CMD ["/server"]