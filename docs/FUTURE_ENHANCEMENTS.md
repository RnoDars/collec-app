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

---

**Note :** Ces améliorations peuvent être planifiées entre deux versions majeures selon les priorités et le temps disponible.

**Dernière mise à jour :** 30/12/2025
