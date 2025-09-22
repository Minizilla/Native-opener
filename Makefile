# Makefile pour microzilla
# Compilation multi-plateforme

# Nom du binaire
BINARY_NAME=microzilla

# Cible par défaut - compile pour toutes les plateformes
all: linux windows darwin darwin-arm64

# Compilation pour Linux (AMD64)
linux:
	@echo "Compilation pour Linux AMD64..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux .

# Compilation pour Windows (AMD64)
windows:
	@echo "Compilation pour Windows AMD64..."
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe .

# Compilation pour macOS Intel (AMD64)
darwin:
	@echo "Compilation pour macOS Intel..."
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin .

# Compilation pour macOS ARM (M1/M2)
darwin-arm64:
	@echo "Compilation pour macOS ARM64..."
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-darwin-arm64 .

# Nettoyage des binaires compilés
clean:
	@echo "Nettoyage des binaires..."
	rm -f $(BINARY_NAME)-linux $(BINARY_NAME).exe $(BINARY_NAME)-darwin $(BINARY_NAME)-darwin-arm64

# Installation des dépendances
deps:
	@echo "Installation des dépendances..."
	go mod download
	go mod tidy

# Test du code
test:
	@echo "Exécution des tests..."
	go test ./...

# Formatage du code
fmt:
	@echo "Formatage du code..."
	go fmt ./...

# Vérification du code
vet:
	@echo "Vérification du code..."
	go vet ./...

# Compilation pour la plateforme actuelle uniquement
build:
	@echo "Compilation pour la plateforme actuelle..."
	go build -o $(BINARY_NAME) .

# Aide
help:
	@echo "Commandes disponibles:"
	@echo "  all          - Compile pour toutes les plateformes"
	@echo "  linux        - Compile pour Linux AMD64"
	@echo "  windows      - Compile pour Windows AMD64"
	@echo "  darwin       - Compile pour macOS Intel"
	@echo "  darwin-arm64 - Compile pour macOS ARM64"
	@echo "  build        - Compile pour la plateforme actuelle"
	@echo "  clean        - Supprime tous les binaires"
	@echo "  deps         - Installe les dépendances"
	@echo "  test         - Exécute les tests"
	@echo "  fmt          - Formate le code"
	@echo "  vet          - Vérifie le code"
	@echo "  help         - Affiche cette aide"

# Déclaration des cibles qui ne correspondent pas à des fichiers
.PHONY: all linux windows darwin darwin-arm64 clean deps test fmt vet build help
