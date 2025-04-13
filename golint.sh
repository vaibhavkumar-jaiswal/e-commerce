echo "🔄 Cleaning old build..."
go clean

echo "🎨 Formatting code..."
goimports -w .

echo "🔍 Running golangci-lint..."
golangci-lint run ./...