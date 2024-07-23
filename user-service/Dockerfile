FROM golang:1.19.3-alpine3.16
WORKDIR /app

COPY . .

COPY go.mod .
COPY go.sum .
COPY config.yml .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o api .

ENTRYPOINT ["./api"]