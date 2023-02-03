FROM golang:1.19-bullseye
#falta adicionar as envs aqui

WORKDIR /app
COPY . .

RUN go mod tidy

ENTRYPOINT [ "go", "run", "main.go"]