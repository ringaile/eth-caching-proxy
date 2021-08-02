FROM golang:1.15-alpine

WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o /eth-caching-proxy

EXPOSE 9000

CMD [ "/eth-caching-proxy" ]