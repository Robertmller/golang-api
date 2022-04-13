FROM golang:1.16-alpine

WORKDIR /code

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /golang-api

EXPOSE 8000

CMD [ "/golang-api" ]