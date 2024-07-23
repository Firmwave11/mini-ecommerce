FROM golang:latest
WORKDIR /app

COPY . .

COPY go.mod .
COPY go.sum .
COPY go.work .
COPY go.work.sum .
COPY config.yml .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o api .

ENTRYPOINT ["./api"]