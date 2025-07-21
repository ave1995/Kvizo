FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app

RUN go build -o main . \
    && chmod +x main

CMD [ "/app/main" ]