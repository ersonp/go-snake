run: ## Run the code
	go run main.go
	
format: ## Formats the code. Must have goimports installed
	goimports -w -local github.com/ersonp/go-snake ./cmd
	goimports -w -local github.com/ersonp/go-snake ./pkg
