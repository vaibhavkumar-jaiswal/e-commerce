#!/bin/bash

set -e  # exit on error

PORT=8080

echo "🛑 Checking if port $PORT is already in use..."
if lsof -i :$PORT -sTCP:LISTEN -t >/dev/null ; then
  PID=$(lsof -i :$PORT -sTCP:LISTEN -t)
  echo "Killing process on port $PORT (PID: $PID)..."
  kill -9 $PID
fi

echo "🔄 Cleaning old build..."
go clean

echo "🎨 Formatting code..."
goimports -w .

echo "📝 Creating Swagger Docs..."
rm -rf docs
swag init -g main.go --parseDependency=false --parseInternal

# echo "🔍 Running golangci-lint..."
# golangci-lint run ./...

echo "🔨 Building the project..."
go build -o bin/e-commerce .

echo "🚀 Running the project..."
./bin/e-commerce
