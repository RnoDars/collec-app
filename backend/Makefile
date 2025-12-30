.PHONY: help test test-unit test-e2e test-coverage run dev

# Variables
GO=go
GOTEST=$(GO) test
APP_NAME=collec-app-backend

help: ## Afficher l'aide
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Exécuter tous les tests
	$(GOTEST) -v ./...

test-unit: ## Exécuter uniquement les tests unitaires (sans E2E)
	$(GOTEST) -v $$(go list ./... | grep -v /test/e2e)

test-e2e: ## Exécuter les tests end-to-end (nécessite serveur en cours d'exécution)
	@echo "⚠️  Assurez-vous que le serveur est démarré sur http://localhost:8080"
	@echo "   Commande: make run (dans un autre terminal)"
	@echo ""
	$(GOTEST) -v ./test/e2e/...

test-coverage: ## Exécuter les tests avec couverture (unitaires seulement)
	$(GOTEST) -coverprofile=coverage.out $$(go list ./... | grep -v /test/e2e)
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "📊 Rapport de couverture généré : coverage.html"

run: ## Démarrer le serveur
	$(GO) run cmd/api/main.go

dev: ## Démarrer le serveur en mode développement (avec air si installé)
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "⚠️  'air' n'est pas installé. Installation recommandée : go install github.com/air-verse/air@latest"; \
		echo "   Démarrage avec 'go run' à la place..."; \
		$(GO) run cmd/api/main.go; \
	fi

build: ## Compiler l'application
	$(GO) build -o bin/$(APP_NAME) cmd/api/main.go

clean: ## Nettoyer les fichiers générés
	rm -f coverage.out coverage.html
	rm -rf bin/

deps: ## Télécharger les dépendances
	$(GO) mod download
	$(GO) mod tidy

.DEFAULT_GOAL := help
