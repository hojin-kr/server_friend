FROM golang:1.21-alpine AS builder

WORKDIR /
COPY . ./
RUN go build -o /app
ENTRYPOINT [ "/app" ]
