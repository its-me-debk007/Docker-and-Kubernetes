FROM golang:alpine

LABEL maintainer="Debashish Kundu"

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 3000

RUN go build -o main

CMD ./main