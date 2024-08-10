# === BUILD STAGE === #
FROM golang:1.22-alpine as build

ARG ACCESS_TOKEN

RUN apk add --no-cache git

WORKDIR /build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test -v ./...
RUN go build -ldflags="-w -s" -o npm-cache-proxy

# === RUN STAGE === #
FROM redis:alpine as run

RUN apk update \
        && apk upgrade \
        && apk add --no-cache ca-certificates \
        && update-ca-certificates \
        && rm -rf /var/cache/apk/*
        
COPY --from=build /build/npm-cache-proxy /app/npm-cache-proxy
COPY entrypoint.sh /srv/entrypoint.sh

ENV LISTEN_ADDRESS 0.0.0.0:8080
ENV GIN_MODE release

ENV REDIS_PASSWORD password
ENV REDIS_ADDRESS 127.0.0.1:6379

VOLUME /data
WORKDIR /data

CMD ["sh", "/srv/entrypoint.sh"]
