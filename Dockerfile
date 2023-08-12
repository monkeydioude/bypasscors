FROM golang:alpine

WORKDIR /usr/src/app

COPY . .
RUN go build -v -o /usr/local/bin/bypasscors ./...

EXPOSE 8080

CMD ["bypasscors"]
