FROM golang:latest

LABEL maintainer = "Niva <niva.ranavat@gmail.com>"

WORKDIR /app 

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8088

RUN go build 

CMD ["./products-api"]