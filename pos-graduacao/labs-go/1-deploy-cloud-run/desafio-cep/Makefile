.PHONY: build run test docker-build docker-up docker-down clean help

# Variáveis
APP_NAME=desafio-cep
DOCKER_IMAGE=desafio-cep:latest

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Compila a aplicação
	@echo "Compilando..."
	@go build -o $(APP_NAME) ./cmd/main.go
	@echo "Build concluído: ./$(APP_NAME)"

run: ## Executa a aplicação localmente
	@echo "Executando aplicação..."
	@go run ./cmd/main.go

test: ## Executa os testes
	@echo "Executando testes..."
	@go test ./... -v

test-coverage: ## Executa os testes com coverage
	@echo "Executando testes com coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report gerado: coverage.html"

docker-build: ## Build da imagem Docker
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .
	@echo "Docker image criada: $(DOCKER_IMAGE)"

docker-up: ## Sobe o container com docker-compose
	@echo "Iniciando serviços..."
	@docker compose up -d --build
	@echo "Serviços iniciados na porta 8080"

docker-down: ## Para os containers
	@echo "Parando serviços..."
	@docker compose down
	@echo "Serviços parados"

docker-logs: ## Mostra os logs dos containers
	@docker compose logs -f

clean: ## Remove arquivos gerados
	@echo "Limpando..."
	@rm -f $(APP_NAME)
	@rm -f coverage.out coverage.html
	@echo "Arquivos limpos"

.DEFAULT_GOAL := help
