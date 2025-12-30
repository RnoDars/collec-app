# Problèmes connus - v0.1.0

## Problèmes identifiés lors des tests Docker (30/12/2025)

### 1. Services qui fonctionnent ✅
- **PostgreSQL** - Démarré et opérationnel (port 5432)
- **Kafka + Zookeeper** - Démarrés et opérationnels (port 9092)
- **Prometheus** - Démarré après correction permissions (port 9090)
- **Grafana** - Démarré après correction permissions (port 3001)
- **Promtail** - Démarré

### 2. Problèmes de sécurité 🔒

#### 2.1 Mot de passe Grafana faible ❌ CORRIGÉ
**Symptôme :** Le mot de passe admin de Grafana était en dur dans docker-compose.yml avec une valeur faible (`admin/admin`)

**Risque :** Accès non autorisé à Grafana et aux données de monitoring en production

**Solution appliquée :**
- Utilisation de variables d'environnement : `${GRAFANA_ADMIN_USER:-admin}` et `${GRAFANA_ADMIN_PASSWORD:-changeme}`
- Création d'un fichier `.env.example` à la racine du projet
- Documentation claire pour changer ces valeurs en production

**Fichiers modifiés :**
- `docker-compose.yml` - Variables d'environnement au lieu de valeurs en dur
- `.env.example` - Template avec instructions

**Action utilisateur requise :**
```bash
# Créer un fichier .env avec des mots de passe sécurisés
cp .env.example .env
# Éditer .env et changer GRAFANA_ADMIN_PASSWORD
```

**Status :** ✅ Corrigé

#### 2.2 Mot de passe PostgreSQL faible ⚠️
**Note :** Le mot de passe PostgreSQL est également faible (`postgres/postgres`) mais restera en dur pour le développement local. **À CHANGER ABSOLUMENT EN PRODUCTION** via variables d'environnement.

#### 2.3 Problèmes de permissions ⚠️
**Symptôme :** Les fichiers de configuration dans `monitoring/` avaient des permissions restrictives

**Solution appliquée :**
```bash
chmod -R 755 monitoring/
docker compose restart prometheus grafana loki
```

**Status :** Résolu pour Prometheus et Grafana

### 3. Configuration Loki obsolète ✅ CORRIGÉ
**Symptôme :** Loki ne démarrait pas avec plusieurs erreurs de configuration

**Erreurs identifiées :**
```
- Schema v11 utilisé au lieu de v13 (requis pour Structured Metadata)
- Index type `boltdb-shipper` au lieu de `tsdb`
```

**Fichier concerné :** `monitoring/loki/loki-config.yml`

**Solution appliquée :**
1. ✅ Mise à jour du schema_config vers v13
2. ✅ Changement de l'index type de `boltdb-shipper` vers `tsdb`
3. ✅ Redémarrage du service Loki

**Status :** ✅ Corrigé - Loki fonctionne maintenant correctement

### 4. Warning docker-compose ✅ CORRIGÉ
**Symptôme :** `the attribute 'version' is obsolete`

**Fichier concerné :** `docker-compose.yml` (ligne 1)

**Solution appliquée :** Suppression de la ligne `version: '3.8'` (obsolète en Docker Compose v2)

**Status :** ✅ Corrigé

## Tâches à réaliser pour finaliser v0.1.0

- [x] **Corriger la configuration Loki** ✅ TERMINÉ
  - Mise à jour vers schema v13
  - Changement vers index type tsdb
  - Test du démarrage réussi

- [x] **Supprimer la ligne version dans docker-compose.yml** ✅ TERMINÉ

- [x] **Tester l'infrastructure complète** ✅ TERMINÉ
  - Tous les services démarrés avec succès
  - PostgreSQL, Kafka, Zookeeper, Prometheus, Grafana, Loki, Promtail opérationnels

- [ ] **Documenter l'installation Docker** (priorité moyenne)
  - Ajouter instructions détaillées sur les permissions
  - Ajouter procédure de vérification des services

- [ ] **Créer des health checks** (priorité basse)
  - Ajouter health checks pour tous les services
  - Script de vérification automatique

- [ ] **Tests end-to-end complets** (priorité moyenne)
  - Tester les connexions entre services
  - Accéder à Grafana et configurer les dashboards

## Services testés avec succès

### Backend Go ✅
```bash
cd backend
go run cmd/api/main.go
# Output: Collec-App Backend v0.1.0
# Server ready to start on port 8080
```

### Frontend Next.js ✅
```bash
cd frontend
npm run dev
# Output: Ready in 429ms
# Local: http://localhost:3000
```

### Services Docker ✅ TOUS OPÉRATIONNELS
- PostgreSQL: ✅ Opérationnel
- Kafka + Zookeeper: ✅ Opérationnel
- Prometheus: ✅ Opérationnel (après correction permissions)
- Grafana: ✅ Opérationnel (après correction permissions + sécurité)
- Loki: ✅ Opérationnel (après correction configuration)
- Promtail: ✅ Opérationnel

## Prochaines étapes

1. ✅ ~~Corriger la configuration Loki~~ - TERMINÉ
2. ✅ ~~Tester l'infrastructure complète~~ - TERMINÉ
3. 📝 Finaliser la documentation (health checks, guides détaillés)
4. 🚀 Passer à la v0.2.0 (Authentification)

---

**Dernière mise à jour :** 30/12/2025
**Status global v0.1.0 :** ✅ Infrastructure complète et fonctionnelle
