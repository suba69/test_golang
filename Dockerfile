FROM golang:1.22-alpine

WORKDIR /app

COPY . /app
RUN go build -o ./test_golang/main.go
CMD [ "./test_golang" ]
