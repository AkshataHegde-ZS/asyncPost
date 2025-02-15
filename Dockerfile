FROM golang:latest

WORKDIR /app

COPY ./configs /configs
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000
CMD ["./main"]