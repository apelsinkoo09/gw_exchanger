FROM golang:1.23.5-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# Копирование зависимостей в образ
COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/exchanger ./cmd/main.go

FROM alpine 

COPY --from=builder /usr/local/src/bin/exchanger /
COPY /config.env /config.env


CMD ["/exchanger"]