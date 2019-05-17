init:
	cp -r ./docker/.env.example ./docker/.env

linter:
	golangci-lint run ./...

leak:
	go build -gcflags "-m -l" main.go

leak-more:
	go build -gcflags "-m -l -S" main.go
