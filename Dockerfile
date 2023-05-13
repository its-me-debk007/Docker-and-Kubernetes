FROM golang:alpine

LABEL maintainer="Debashish Kundu"

WORKDIR /app

COPY . .

RUN go mod download

ENV PORT=3000
ENV DATABASE_URL=postgres://postgres:12345@database:5432/docker_and_kubernetes

EXPOSE 3000

RUN go build -o main

CMD ./main