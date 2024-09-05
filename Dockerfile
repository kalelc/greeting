# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build main.go

EXPOSE 8080

# Run
CMD ["./server"]
