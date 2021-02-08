FROM golang:1.15

WORKDIR /go/src/app
COPY . .

CMD ["go", "run", "cmd/app/main.go"]