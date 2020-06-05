FROM golang:1.14.4-alpine3.12

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 9000

CMD ["/app/main"]