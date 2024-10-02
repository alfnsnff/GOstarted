FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /GOstarted

ENV DB_HOST=35.198.196.14
ENV DB_USER=postgres
ENV DB_PASSWORD=umbresgurgur
ENV DB_NAME=gostarted
ENV DB_PORT=5432
ENV DB_SSLMODE=disable

EXPOSE 8080

CMD ["/GOstarted"]
