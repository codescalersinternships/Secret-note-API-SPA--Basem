FROM golang:alpine3.20 AS build

WORKDIR /app

RUN apk add --no-cache \
    gcc \
    musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY  . . 

RUN CGO_ENABLED=1 GOOS=linux go build -o /goapp ./cmd

FROM alpine:edge

WORKDIR /app
COPY .env .env

COPY --from=build /goapp .

EXPOSE ${PORT}

ENTRYPOINT ["./goapp"]

