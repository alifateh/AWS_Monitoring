FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go mod init alerting-app || true
RUN go mod tidy || true

EXPOSE 8080

CMD ["go", "run", "main.go"]