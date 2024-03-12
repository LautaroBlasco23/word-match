FROM golang:1.22

WORKDIR /app

COPY . .

ENV REDIS_PASSWORD=test123

RUN go build

EXPOSE 8002

CMD go run main.go