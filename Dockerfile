FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /GOstarted

EXPOSE 8080

CMD ["/GOstarted"]
