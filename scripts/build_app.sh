#!/bin/sh

echo "Generating Swagger documentation"
swag init -g cmd/main.go

echo "Building"
go build -o service_executable ./cmd