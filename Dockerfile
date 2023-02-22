FROM golang:1.19.4-alpine
RUN mkdir /app
ADD go.mod go.sum . /app/

WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]