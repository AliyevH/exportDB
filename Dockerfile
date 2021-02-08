FROM golang:1.15

WORKDIR /app
COPY . .

CMD ["go", "run", "cmd/app/main.go"]