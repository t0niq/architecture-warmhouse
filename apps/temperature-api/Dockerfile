FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /temperature-api

EXPOSE 8081

CMD ["/temperature-api"] 