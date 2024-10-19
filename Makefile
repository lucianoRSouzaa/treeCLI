.PHONY: build run clean help

BINARY_NAME=treecli
BUILD_DIR=bin

build:
	@echo "Compilando o projeto..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/treecli/main.go
	@echo "Pronto! Executável gerado: $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Limpando arquivos..."
	@rm -rf $(BUILD_DIR)
	@echo "Arquivos removidos."

help:
	@echo "Comandos disponíveis:"
	@echo "  make build               - Compila o projeto"
	@echo "  make clean               - Remove o executável e arquivos temporários"
	@echo "  make help                - Exibe esta mensagem de ajuda"
