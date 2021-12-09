FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .
CMD ["go", "run", "cmd/api/main.go"]