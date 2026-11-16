
FROM golang:1.27-alpine

WORKDIR /app

# dependency files
COPY go.mod go.sum ./

RUN go mod download

# copy application code
COPY ./main.go .

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
