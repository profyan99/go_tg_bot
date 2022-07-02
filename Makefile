.PHONY: run
run:
	go run cmd/bot/main.go

lint:  ## Run lint the current project
	golangci-lint run --fix