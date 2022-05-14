FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o task-api

EXPOSE 8080

CMD ./task-api
