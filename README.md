# Collec-App 📦

Application web moderne de gestion de collections avec GoLang et Next.js.

## 🎯 Vision

Créer une application performante permettant aux utilisateurs de gérer, organiser et partager leurs collections d'objets de manière intuitive et collaborative.

## ✨ Fonctionnalités (Roadmap)

### v0.1.0 - Configuration initiale ✅ (Infrastructure complète)
- [x] Configuration complète du projet
- [x] Architecture monorepo (backend Go + frontend Next.js)
- [x] Stack de monitoring (Prometheus, Grafana, Loki)
- [x] Infrastructure Docker (PostgreSQL, Kafka)
- [x] Tests Backend Go ✅
- [x] Tests Frontend Next.js ✅
- [x] Tests Infrastructure Docker ✅
  - ✅ PostgreSQL opérationnel
  - ✅ Kafka + Zookeeper opérationnels
  - ✅ Prometheus opérationnel
  - ✅ Grafana opérationnel (sécurisé avec variables d'environnement)
  - ✅ Loki opérationnel (configuration v13 + tsdb)
  - ✅ Promtail opérationnel
- [x] Documentation détaillée installation Docker ✅
- [x] Tests end-to-end complets ✅

**📋 Voir [KNOWN_ISSUES.md](docs/KNOWN_ISSUES.md) pour l'historique des corrections**

### v0.2.0 - Authentification ✅ (Complète et fonctionnelle)
- [x] Backend : Inscription et connexion utilisateur
- [x] Backend : Système JWT avec refresh tokens
- [x] Backend : Endpoints protégés avec middleware
- [x] Backend : Tests unitaires (11 tests, 82.1% couverture service auth)
- [x] Backend : Tests E2E (13 tests automatisés)
- [x] Frontend : Composants LoginForm et RegisterForm
- [x] Frontend : Pages d'authentification (/login, /register, /profile)
- [x] Frontend : Store Zustand avec persistance
- [x] Frontend : Tests unitaires (25 tests, 3 suites complètes)
- [x] Tests manuels : Flux complet vérifié
- [x] Méthodologie TDD : Appliquée et documentée

**📋 Voir [V0.2.0_PLAN.md](docs/V0.2.0_PLAN.md) pour le plan détaillé et [TDD_WORKFLOW.md](docs/TDD_WORKFLOW.md) pour la méthodologie**

**🎯 Prochaine étape prioritaire :** UI/UX - Amélioration du design (voir [FUTURE_ENHANCEMENTS.md](docs/FUTURE_ENHANCEMENTS.md))

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

4. **Configuration Docker (IMPORTANT pour la sécurité)**
```bash
# À la racine du projet
cp .env.example .env
# Éditer .env et changer GRAFANA_ADMIN_PASSWORD avec un mot de passe fort !
```

5. **Lancer l'infrastructure avec Docker**
```bash
docker compose up -d
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

- **Grafana:** http://localhost:3001 (voir `.env` pour les identifiants)
- **Prometheus:** http://localhost:9090
- **Loki:** http://localhost:3100

⚠️ **Sécurité :** Changez le mot de passe Grafana par défaut dans `.env` avant de démarrer Docker !

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

**Version actuelle:** 0.2.0
**Status:** En développement actif
**Dernière mise à jour:** 30/12/2025
