FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]