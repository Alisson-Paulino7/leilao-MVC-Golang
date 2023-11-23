FROM golang:latest

WORKDIR /

COPY ./go.mod .
COPY ./go.sum .

COPY . .

RUN go build -o main .

CMD ["./main"]

