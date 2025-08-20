FROM golang:1.24-alpine

RUN apk add --no-cache git tzdata

ENV TZ=Asia/Bangkok

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV AIR_CONF=.air.toml

CMD ["air", "-c", ".air.toml"]