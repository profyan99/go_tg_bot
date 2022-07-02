.PHONY: run
run:
	go run cmd/bot/main.go

lint:  ## Run lint the current project
	golangci-lint run --fix

generate_orm: ## Generating ent files
	-go run entgo.io/ent/cmd/ent@latest generate --target ./internal/repository/ent ./internal/repository/schema