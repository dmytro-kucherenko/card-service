build:
	@go build -o bin/api cmd/api/main.go

clean:
	@rm -f bin/api

start:
	@bin/api

run:
	@go run cmd/api/main.go

watch:
	@air -c api.air.toml

test:
	@go test ./...

lint:
	@go vet ./...

pre-commit:
	@pre-commit autoupdate && pre-commit install
