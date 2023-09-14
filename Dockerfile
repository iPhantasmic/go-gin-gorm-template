FROM golang:alpine

WORKDIR /app

RUN apk add --no-cache git

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

COPY . .

RUN chmod +x /app/scripts/build_app.sh

CMD ["CompileDaemon", "-build", "./scripts/build_app.sh", "-command", "./service_executable", "-polling=true", "-exclude-dir=docs"]