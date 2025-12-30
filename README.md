# Collec-App 📦

Application web moderne de gestion de collections avec GoLang et Next.js.

## 🎯 Vision

Créer une application performante permettant aux utilisateurs de gérer, organiser et partager leurs collections d'objets de manière intuitive et collaborative.

## ✨ Fonctionnalités (Roadmap)

### v0.1.0 - Configuration initiale 🚧 (En cours de test)
- [x] Configuration complète du projet
- [x] Architecture monorepo (backend Go + frontend Next.js)
- [x] Stack de monitoring (Prometheus, Grafana, Loki)
- [x] Infrastructure Docker (PostgreSQL, Kafka)
- [x] Tests Backend Go ✅
- [x] Tests Frontend Next.js ✅
- [ ] Tests Infrastructure Docker (en attente d'environnement Docker)
- [ ] Tests end-to-end complets

### v0.2.0 - Authentification (Planifié)
- [ ] Inscription et connexion utilisateur
- [ ] Système JWT avec refresh tokens
- [ ] Gestion du profil utilisateur

### v0.3.0+ - Fonctionnalités métier (Planifié)
- [ ] Gestion des collections (CRUD)
- [ ] Gestion des items avec métadonnées
- [ ] Catégories et tags
- [ ] Recherche et filtrage avancés
- [ ] Statistiques et visualisations

## 🏗️ Architecture

**Stack Technique:**
- **Backend:** GoLang 1.21+ avec Gin/Echo
- **Frontend:** Next.js 14 (App Router) + TypeScript + TailwindCSS
- **Base de données:** PostgreSQL 15+
- **Message Queue:** Apache Kafka
- **Monitoring:** Prometheus + Grafana + Loki
- **Containerisation:** Docker + Docker Compose

**Structure Monorepo:**
```
collec-app/
├── backend/          # API Go
├── frontend/         # Application Next.js
├── shared/           # Types partagés
├── monitoring/       # Configuration monitoring
└── docs/            # Documentation
```

## 🚀 Démarrage rapide

### Prérequis

- Go 1.21+
- Node.js 18+
- Docker & Docker Compose
- Git

### Installation

1. **Cloner le repository**
```bash
git clone https://github.com/RnoDars/collec-app.git
cd collec-app
```

2. **Configuration Backend**
```bash
cd backend
cp .env.example .env
# Éditer .env avec vos paramètres
go mod download
```

3. **Configuration Frontend**
```bash
cd frontend
npm install
```

4. **Lancer l'infrastructure avec Docker**
```bash
docker-compose up -d
```

Cela démarre :
- PostgreSQL (port 5432)
- Kafka + Zookeeper (port 9092)
- Prometheus (port 9090)
- Grafana (port 3001) - admin/admin
- Loki (port 3100)

### Développement

**Backend:**
```bash
cd backend
go run cmd/api/main.go
```

**Frontend:**
```bash
cd frontend
npm run dev
```

L'application sera accessible sur http://localhost:3000

## 📊 Monitoring

- **Grafana:** http://localhost:3001 (admin/admin)
- **Prometheus:** http://localhost:9090
- **Loki:** http://localhost:3100

## 🧪 Tests

**Backend:**
```bash
cd backend
go test ./... -v -cover
```

**Frontend:**
```bash
cd frontend
npm test
```

## 📝 Conventions de développement

### Commits
- Format: Conventional Commits (feat:, fix:, docs:, chore:)
- Commits petits et atomiques (toutes les 15-30 min)
- Messages en français et descriptifs

### Tests
- Tests unitaires OBLIGATOIRES pour tout nouveau code
- Couverture minimum: 80%
- Tests critiques: 100%

### Code Style
- Backend: gofmt, golangci-lint
- Frontend: ESLint, Prettier
- Indentation: 2 espaces

## 📚 Documentation

- **Configuration projet:** `.claude-project.json`
- **API Documentation:** À venir (Swagger/OpenAPI)
- **Architecture:** `docs/architecture/`

## 🔒 Sécurité

- Authentification JWT avec refresh tokens
- RBAC (Role-Based Access Control)
- Rate limiting et CORS configurés
- Validation stricte des entrées
- TLS 1.3 en production

## 🤝 Contribution

Ce projet suit une méthodologie itérative :
1. **Besoin** - Définir ou revoir le besoin fonctionnel
2. **Planification** - Créer les tâches et estimer
3. **Développement** - Coder avec tests
4. **Review** - Vérifier la conformité
5. **Itération** - Passer à la suite

## 📄 Licence

Ce projet est privé.

## 👥 Auteurs

- Arnaud Dars - [@RnoDars](https://github.com/RnoDars)

---

**Version actuelle:** 0.1.0
**Status:** En développement actif
