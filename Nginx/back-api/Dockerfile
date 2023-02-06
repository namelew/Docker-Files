FROM golang:1.19-bullseye

ENV DBUSER=postgres
ENV DBHOST=localhost
ENV DBPORT=5432
ENV DBPW=postgres
ENV DBNAME=commerce

WORKDIR /app
COPY . .

RUN go mod tidy

ENTRYPOINT [ "go", "run", "main.go"]