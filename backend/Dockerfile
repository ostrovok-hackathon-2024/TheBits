FROM golang:1.23 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o application ./cmd

FROM ubuntu:22.04

COPY --from=build /app/application /application

ENTRYPOINT exec /application

EXPOSE 8080
EXPOSE 8088
