init:
	cp -r ./docker/.env.example ./docker/.env
linter:
	golangci-lint run ./...
