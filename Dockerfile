FROM golang:1.21-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0  \
    GOARCH="amd64" \
    GOOS=linux

WORKDIR /build
COPY . .
RUN go mod download
WORKDIR /build/cmd/main
RUN go build -o main .

FROM alpine:3.17

RUN apk add --no-cache tzdata
ENV TZ="Asia/Jakarta"

WORKDIR /www

COPY --from=builder /build/cmd/main/main /www/

ENV APP_PORT=
ENV APP_SECRET=

ENV DB_HOST=
ENV DB_USER=
ENV DB_PASS=
ENV DB_NAME=
ENV DB_PORT=
ENV DB_SSL_MODE=
ENV DB_TIMEZONE=

EXPOSE 8000

ENTRYPOINT ["/www/main"]