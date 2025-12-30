# Tests End-to-End (E2E)

Ce dossier contient les tests end-to-end qui vérifient le fonctionnement complet de l'API en conditions réelles.

## Prérequis

Les tests E2E nécessitent que le serveur backend soit démarré et accessible sur `http://localhost:8080`.

### Démarrer l'infrastructure

```bash
# Depuis la racine du projet
sg docker -c "docker compose up -d"
```

### Démarrer le serveur backend

```bash
# Dans un terminal séparé
cd backend
make run
# ou directement : go run cmd/api/main.go
```

## Exécution des tests

### Exécuter tous les tests E2E

```bash
cd backend
make test-e2e
```

### Exécuter un test spécifique

```bash
cd backend
go test -v ./test/e2e -run TestE2E_AuthenticationFlow
```

### Exécuter uniquement les tests unitaires (sans E2E)

```bash
cd backend
make test-unit
```

## Tests disponibles

### `auth_test.go`

**TestE2E_AuthenticationFlow** - Flux complet d'authentification :
1. ✅ Inscription d'un nouvel utilisateur (201)
2. ✅ Inscription avec email existant retourne 409
3. ✅ Connexion avec identifiants valides (200)
4. ✅ Connexion avec mauvais mot de passe (401)
5. ✅ Connexion avec email inexistant (401)
6. ✅ Récupération d'un token valide
7. ✅ Accès à endpoint protégé avec token valide (200)
8. ✅ Accès à endpoint protégé sans token (401)
9. ✅ Accès à endpoint protégé avec token invalide (401)

**TestE2E_ValidationErrors** - Erreurs de validation :
- ✅ Inscription sans email (400)
- ✅ Inscription avec mot de passe trop court (400)
- ✅ Inscription avec email invalide (400)

**TestE2E_HealthCheck** - Vérification de santé :
- ✅ Endpoint `/health` retourne 200

## Architecture des tests E2E

```
backend/
└── test/
    └── e2e/
        ├── README.md          # Ce fichier
        └── auth_test.go       # Tests d'authentification
```

## Bonnes pratiques

### 1. Email unique par test
Les tests génèrent des emails uniques avec un timestamp pour éviter les conflits :
```go
timestamp := time.Now().Unix()
testEmail := fmt.Sprintf("e2e-test-%d@example.com", timestamp)
```

### 2. Tests indépendants
Chaque test doit pouvoir s'exécuter indépendamment des autres.

### 3. Nettoyage
Les tests utilisent la base de données réelle. Considérez d'ajouter un nettoyage si nécessaire :
- Option 1 : Base de données de test séparée
- Option 2 : Nettoyage manuel après les tests
- Option 3 : Utiliser des transactions rollback

### 4. CI/CD
Pour l'intégration continue, démarrez les services avant les tests :

```yaml
# Exemple GitHub Actions
steps:
  - name: Start services
    run: docker compose up -d

  - name: Wait for backend
    run: |
      timeout 30 bash -c 'until curl -f http://localhost:8080/health; do sleep 1; done'

  - name: Run E2E tests
    run: cd backend && make test-e2e
```

## Dépendances

Les tests utilisent :
- `github.com/stretchr/testify` : assertions et require
- `net/http` : requêtes HTTP
- `encoding/json` : sérialisation/désérialisation JSON

## Ajout de nouveaux tests E2E

Pour ajouter de nouveaux tests :

1. Créer un nouveau fichier `*_test.go` dans `test/e2e/`
2. Utiliser le package `e2e_test`
3. Préfixer les tests avec `TestE2E_`
4. S'assurer que le serveur est démarré avant les tests

Exemple :

```go
package e2e_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestE2E_MonNouveauTest(t *testing.T) {
	t.Run("description du test", func(t *testing.T) {
		// Code du test
		assert.True(t, true)
	})
}
```

## Troubleshooting

### Erreur "connection refused"
Le serveur n'est pas démarré. Lancez `make run` dans un terminal séparé.

### Tests échouent avec 409 (email existant)
Les emails de test existent déjà en base. Les tests génèrent des emails uniques avec timestamp pour éviter ce problème.

### Base de données non accessible
Vérifiez que PostgreSQL est démarré :
```bash
sg docker -c "docker compose ps"
```

## Prochaines améliorations

- [ ] Tests E2E pour les collections (v0.3.0)
- [ ] Tests E2E pour les items (v0.4.0)
- [ ] Base de données de test séparée
- [ ] Script de cleanup automatique
- [ ] Tests de performance/charge
