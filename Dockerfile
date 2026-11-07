FROM golang:1.26.8-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["make", "server-run"]