FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /Microservice_TS_1

EXPOSE 8080

CMD ["/Microservice_TS_1"]
