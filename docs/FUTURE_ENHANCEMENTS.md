# Améliorations Futures

Ce fichier liste les fonctionnalités et améliorations identifiées mais reportées à des versions ultérieures.

## 🔐 Authentification & Profil

### Priorité Haute
- [ ] **Modification du profil utilisateur**
  - Changer l'email
  - Changer le mot de passe
  - Version cible : v0.2.1 ou v0.3.0

### Priorité Moyenne
- [ ] **Remember Me**
  - Checkbox "Se souvenir de moi" sur la page de connexion
  - Prolonge la durée du refresh token (30 jours au lieu de 7)
  - Version cible : v0.2.1 ou future

- [ ] **Vérification email**
  - Envoi d'un email de confirmation lors de l'inscription
  - Lien de validation avant de pouvoir se connecter
  - Nécessite : Configuration SMTP
  - Version cible : v0.3.0 ou future

- [ ] **Récupération de mot de passe**
  - "Mot de passe oublié ?" sur la page de connexion
  - Email avec lien de réinitialisation
  - Nécessite : Configuration SMTP
  - Version cible : v0.3.0 ou future

### Priorité Basse
- [ ] **Suppression de compte**
  - Permettre à l'utilisateur de supprimer son compte
  - Confirmation requise + suppression cascade des collections
  - Version cible : v1.0.0 ou future

- [ ] **Authentification multi-facteurs (2FA)**
  - TOTP ou SMS pour sécurité renforcée
  - Version cible : v1.x.x

- [ ] **OAuth / Social login**
  - Connexion via Google, GitHub, etc.
  - Version cible : v1.x.x

## 🎨 UI/UX

### Priorité Haute
- [ ] **Amélioration de l'interface utilisateur**
  - Design system : définir palette de couleurs, typographie, espacements
  - Composants UI réutilisables : buttons, inputs, cards, modals
  - Layout responsive amélioré
  - Animations et transitions
  - Considérer : Tailwind UI, shadcn/ui, ou design custom
  - Version cible : v0.2.1 ou v0.3.0
  - **Note :** UI actuelle fonctionnelle mais basique, à améliorer pour meilleure UX

## 📊 Dashboard

- [ ] **Page Dashboard dédiée**
  - Actuellement : redirection vers homepage après connexion
  - Future : redirection vers /dashboard avec vue d'ensemble des collections
  - Version cible : v0.3.0

## 🚀 Infrastructure

### Priorité Moyenne
- [ ] **Documentation détaillée installation Docker**
  - Instructions complètes sur les permissions
  - Procédures de vérification des services
  - Troubleshooting commun
  - Version cible : v0.2.0 ou v0.3.0

- [ ] **Tests end-to-end complets**
  - Tests E2E entre tous les services
  - Playwright ou Cypress
  - Version cible : v0.3.0 ou future

### Priorité Basse
- [ ] **Scripts de health checks automatisés**
  - Vérification automatique de l'état des services Docker
  - Alertes si un service est down
  - Version cible : v0.3.0 ou future

## 🏗️ Architecture & Scalabilité

### Migration vers Microservices (Post-MVP)

**Contexte actuel (v0.1.0 - v1.0.0) :**
- Architecture : **Monolithe modulaire**
- Structure par domaines : auth, collections, items
- Code découplé et prêt pour migration future

**Migration planifiée (v1.0+) :**

- [ ] **Évaluation des besoins de scalabilité**
  - Analyser les métriques de charge
  - Identifier les goulots d'étranglement
  - Décider quels modules migrer en priorité
  - Version cible : v1.0.0

- [ ] **Découpage en microservices**
  - **Phase 1 : Auth Service indépendant**
    - Service d'authentification isolé
    - Base de données dédiée
    - API Gateway pour routage
    - Version cible : v1.1.0

  - **Phase 2 : Collections Service**
    - Service de gestion des collections
    - Base de données séparée
    - Communication via Kafka
    - Version cible : v1.2.0

  - **Phase 3 : Items Service**
    - Service de gestion des items
    - Recherche avec Elasticsearch (optionnel)
    - Version cible : v1.3.0

- [ ] **Infrastructure microservices**
  - API Gateway (Kong, Traefik, ou custom)
  - Service Discovery (Consul, etcd)
  - Distributed tracing (Jaeger, Zipkin)
  - Centralized logging (ELK Stack)
  - Configuration centralisée (Consul, etcd)
  - Version cible : v1.x.x

- [ ] **Communication inter-services**
  - Kafka pour événements asynchrones
  - gRPC pour communication synchrone
  - Circuit breakers (resilience)
  - Retry policies
  - Version cible : v1.x.x

- [ ] **Orchestration & Déploiement**
  - Kubernetes pour orchestration
  - Helm charts
  - CI/CD par service
  - Blue/Green deployment
  - Version cible : v2.0.0

**Avantages attendus :**
- ✅ Scalabilité indépendante par service
- ✅ Déploiements sans downtime
- ✅ Isolation des pannes
- ✅ Technologies différentes par service si besoin
- ✅ Équipes autonomes par service

**Complexité ajoutée :**
- ⚠️ Debugging distribué plus complexe
- ⚠️ Overhead de communication réseau
- ⚠️ Gestion de la cohérence des données
- ⚠️ Infrastructure plus lourde

**Décision :** Monolithe modulaire jusqu'à v1.0, puis migration progressive basée sur les besoins réels de scalabilité.

---

**Note :** Ces améliorations peuvent être planifiées entre deux versions majeures selon les priorités et le temps disponible.

**Dernière mise à jour :** 30/12/2025
