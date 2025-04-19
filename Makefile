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

openapi-gen:
	@go mod vendor
	@touch api/gen/rest/swagger.json
	@go tool github.com/swaggo/swag/cmd/swag init -o api/gen/rest -d cmd/api,internal
	@go tool github.com/swaggo/swag/cmd/swag fmt

api-gen:
	@make openapi-gen
