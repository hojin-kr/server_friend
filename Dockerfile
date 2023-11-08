FROM golang:1.21-alpine AS builder

WORKDIR /
COPY . ./
RUN go build -o /app

FROM golang:1.21-alpine
COPY --from=builder /app /app
ENTRYPOINT [ "/app" ]
