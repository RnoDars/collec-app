# Problèmes connus - v0.1.0

## Problèmes identifiés lors des tests Docker (30/12/2025)

### 1. Services qui fonctionnent ✅
- **PostgreSQL** - Démarré et opérationnel (port 5432)
- **Kafka + Zookeeper** - Démarrés et opérationnels (port 9092)
- **Prometheus** - Démarré après correction permissions (port 9090)
- **Grafana** - Démarré après correction permissions (port 3001)
- **Promtail** - Démarré

### 2. Problèmes de permissions ⚠️
**Symptôme :** Les fichiers de configuration dans `monitoring/` avaient des permissions restrictives

**Solution appliquée :**
```bash
chmod -R 755 monitoring/
docker compose restart prometheus grafana loki
```

**Status :** Résolu pour Prometheus et Grafana

### 3. Configuration Loki obsolète ❌
**Symptôme :** Loki ne démarre pas avec plusieurs erreurs de configuration

**Erreurs :**
```
- Schema v11 utilisé au lieu de v13 (requis pour Structured Metadata)
- Index type `boltdb-shipper` au lieu de `tsdb`
- Permissions sur /etc/loki/local-config.yaml
```

**Fichier concerné :** `monitoring/loki/loki-config.yml`

**Actions nécessaires :**
1. Mettre à jour le schema_config vers v13
2. Changer l'index type de `boltdb-shipper` vers `tsdb`
3. Ajouter `allow_structured_metadata: false` temporairement OU migrer vers schema v13

### 4. Warning docker-compose ⚠️
**Symptôme :** `the attribute 'version' is obsolete`

**Fichier concerné :** `docker-compose.yml` (ligne 1)

**Action :** Supprimer la ligne `version: '3.8'` (obsolète en Docker Compose v2)

## Tâches à réaliser pour finaliser v0.1.0

- [ ] **Corriger la configuration Loki** (priorité haute)
  - Mettre à jour vers schema v13
  - Changer vers index type tsdb
  - Tester le démarrage

- [ ] **Supprimer la ligne version dans docker-compose.yml** (priorité basse)

- [ ] **Documenter l'installation Docker** (priorité moyenne)
  - Ajouter instructions sur les permissions
  - Ajouter procédure de vérification des services

- [ ] **Créer des health checks** (priorité moyenne)
  - Ajouter health checks pour tous les services
  - Script de vérification automatique

- [ ] **Tester l'infrastructure complète** (priorité haute)
  - Vérifier tous les services démarrés
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

### Services Docker partiellement ✅
- PostgreSQL: ✅ Opérationnel
- Kafka: ✅ Opérationnel
- Prometheus: ✅ Opérationnel (après correction)
- Grafana: ✅ Opérationnel (après correction)
- Loki: ❌ Configuration à corriger
- Promtail: ⚠️ En attente de Loki

## Prochaines étapes

1. Corriger la configuration Loki
2. Tester l'infrastructure complète
3. Finaliser la documentation
4. Passer à la v0.2.0 (Authentification)

---

**Dernière mise à jour :** 30/12/2025
**Status global v0.1.0 :** 🚧 En cours de finalisation
