# Méthodologie TDD (Test-Driven Development)

Ce document décrit la méthodologie **TDD obligatoire** pour tout développement dans le projet Collec-App.

## 🎯 Principe fondamental

> **"Écrire les tests AVANT le code d'implémentation"**

Le TDD inverse l'ordre traditionnel : au lieu de coder puis tester, on **teste puis on code**.

## 🔄 Cycle RED-GREEN-REFACTOR

```
┌─────────────┐
│   🔴 RED    │  1. Écrire un test qui échoue
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  📝 COMMIT  │  2. Commiter le test (prefix: test:)
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  🟢 GREEN   │  3. Écrire le code minimum pour passer le test
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  📝 COMMIT  │  4. Commiter le code (prefix: feat:/fix:)
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ 🔵 REFACTOR │  5. Améliorer le code (optionnel)
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  📝 COMMIT  │  6. Commiter le refactoring (prefix: refactor:)
└──────┬──────┘
       │
       └──────> Retour au début pour la prochaine fonctionnalité
```

## 📋 Étapes détaillées

### Étape 1 : Planification (5-10 min)

**Objectif :** Définir clairement ce qu'on va développer

**Actions :**
- Lire la spécification fonctionnelle
- Identifier les cas d'usage (success + edge cases)
- Lister les tests à écrire

**Exemple :**
```
Fonctionnalité : Inscription utilisateur

Cas à tester :
✓ Inscription réussie avec email et password valides
✓ Email déjà existant → erreur 409
✓ Email invalide → erreur 400
✓ Mot de passe trop court → erreur 400
✓ Champs manquants → erreur 400
```

### Étape 2 : Écrire les tests (RED 🔴)

**Objectif :** Écrire les tests qui définissent le comportement attendu

**Actions :**
1. Créer le fichier de test : `*_test.go` (Go) ou `*.test.tsx` (React)
2. Écrire TOUS les tests pour la fonctionnalité
3. Les tests doivent échouer (code pas encore écrit)
4. Exécuter les tests pour confirmer qu'ils échouent

**Backend Go exemple :**
```go
// internal/service/user_service_test.go
func TestRegisterUser(t *testing.T) {
    t.Run("devrait créer un utilisateur avec email valide", func(t *testing.T) {
        // Arrange
        service := NewUserService(mockRepo)
        req := RegisterRequest{Email: "test@example.com", Password: "password123"}

        // Act
        user, err := service.Register(req)

        // Assert
        assert.NoError(t, err)
        assert.Equal(t, "test@example.com", user.Email)
    })

    t.Run("devrait retourner erreur si email existe déjà", func(t *testing.T) {
        // Test qui échoue car la fonction n'existe pas encore
        // ...
    })
}
```

**Frontend React exemple :**
```tsx
// src/components/RegisterForm.test.tsx
describe('RegisterForm', () => {
  it('devrait afficher le formulaire d\'inscription', () => {
    render(<RegisterForm />);
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    // Test qui échoue car le composant n'existe pas encore
  });
});
```

**Commandes :**
```bash
# Backend
cd backend && go test ./... -v

# Frontend
cd frontend && npm test

# Les tests doivent ÉCHOUER ❌
```

### Étape 3 : Commit des tests (📝)

**Objectif :** Sauvegarder les tests avant d'écrire le code

**Format du commit :**
```
test: ajoute tests pour [fonctionnalité]

Tests ajoutés:
- [fichier_test.go/tsx] (X tests)
- Couvre: [cas 1], [cas 2], [cas 3]

❌ Status: Tests échouent (RED) - implémentation à venir
```

**Exemple :**
```bash
git add internal/service/user_service_test.go
git commit -m "test: ajoute tests pour l'inscription utilisateur

Tests ajoutés:
- user_service_test.go (5 tests)
- Couvre: inscription réussie, email existant, validations

❌ Status: 0/5 tests passent (RED) - implémentation à venir
"
```

**Pourquoi commiter des tests qui échouent ?**
- ✅ Historique clair : on voit d'abord les spécifications (tests)
- ✅ Séparation des responsabilités : tests vs implémentation
- ✅ Facilite la revue de code : on comprend l'intention avant le code
- ✅ Impossible d'oublier d'écrire les tests

### Étape 4 : Écrire le code (GREEN 🟢)

**Objectif :** Écrire le code minimum pour faire passer les tests

**Principe YAGNI (You Aren't Gonna Need It) :**
- N'implémenter QUE ce qui est nécessaire pour les tests
- Pas de fonctionnalités "au cas où"
- Pas de sur-ingénierie

**Actions :**
1. Créer les fichiers source
2. Implémenter le strict minimum pour passer les tests
3. Exécuter les tests → tous doivent passer ✅
4. Vérifier la couverture (minimum 80%)

**Backend Go exemple :**
```go
// internal/service/user_service.go
type UserService struct {
    repo UserRepository
}

func (s *UserService) Register(req RegisterRequest) (*User, error) {
    // Validation email
    if !isValidEmail(req.Email) {
        return nil, ErrInvalidEmail
    }

    // Vérifier si email existe
    exists, err := s.repo.EmailExists(req.Email)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, ErrEmailAlreadyExists
    }

    // Créer l'utilisateur
    user := &User{
        ID:    uuid.New().String(),
        Email: req.Email,
    }

    return s.repo.Create(user)
}
```

**Commandes :**
```bash
# Backend
cd backend && go test ./... -v
# ✅ Tous les tests doivent PASSER

# Vérifier la couverture
cd backend && go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Frontend
cd frontend && npm test
cd frontend && npm run test:coverage
```

### Étape 5 : Commit du code (📝)

**Format du commit :**
```
feat: implémente [fonctionnalité]

Code:
- [fichier1.go/tsx]
- [fichier2.go/tsx]

✅ Tests: X/X passent (GREEN)
✅ Couverture: Y%

Tests manuels:
✅ [test manuel 1]
✅ [test manuel 2]
```

**Exemple :**
```bash
git add internal/service/user_service.go
git commit -m "feat: implémente l'inscription utilisateur

Code:
- user_service.go
- user_repository.go

✅ Tests: 5/5 passent (GREEN)
✅ Couverture: 92%

Tests manuels:
✅ curl POST /api/auth/register (success)
✅ curl POST /api/auth/register (email exists → 409)
"
```

### Étape 6 : Refactoring (REFACTOR 🔵) - Optionnel

**Objectif :** Améliorer le code sans changer le comportement

**Quand refactorer ?**
- Code dupliqué
- Fonctions trop longues
- Noms de variables peu clairs
- Structure complexe

**Actions :**
1. Améliorer le code
2. Exécuter les tests → doivent TOUJOURS passer ✅
3. Commiter si les améliorations sont significatives

**Format du commit :**
```
refactor: améliore [aspect] de [fonctionnalité]

- Amélioration 1
- Amélioration 2

✅ Tests: X/X passent (toujours GREEN)
```

## 📊 Exemples complets

### Exemple Backend : Endpoint de Login

#### 1. Tests (RED 🔴)

```go
// internal/handler/auth_handler_test.go
func TestLoginHandler(t *testing.T) {
    t.Run("devrait retourner 200 avec token JWT valide", func(t *testing.T) {
        // Setup
        mockService := new(MockAuthService)
        handler := NewAuthHandler(mockService)

        // Mock
        mockService.On("Login", mock.Anything).Return(&AuthResponse{
            AccessToken: "valid-token",
            User: User{ID: "123", Email: "test@example.com"},
        }, nil)

        // Request
        req := httptest.NewRequest("POST", "/api/auth/login", strings.NewReader(`{
            "email": "test@example.com",
            "password": "password123"
        }`))

        // Act
        resp := handler.Login(req)

        // Assert
        assert.Equal(t, 200, resp.StatusCode)
        assert.NotEmpty(t, resp.Body.AccessToken)
    })

    t.Run("devrait retourner 401 si credentials invalides", func(t *testing.T) {
        // ...
    })
}
```

**Commit :**
```bash
git commit -m "test: ajoute tests pour le login endpoint

Tests ajoutés:
- auth_handler_test.go (4 tests)
- Couvre: login success, invalid credentials, missing fields, service errors

❌ Status: 0/4 tests passent (RED)
"
```

#### 2. Implémentation (GREEN 🟢)

```go
// internal/handler/auth_handler.go
func (h *AuthHandler) Login(req *http.Request) *Response {
    var loginReq LoginRequest
    if err := json.NewDecoder(req.Body).Decode(&loginReq); err != nil {
        return ErrorResponse(400, "Invalid request body")
    }

    authResp, err := h.service.Login(loginReq)
    if err != nil {
        if errors.Is(err, ErrInvalidCredentials) {
            return ErrorResponse(401, "Invalid credentials")
        }
        return ErrorResponse(500, "Internal server error")
    }

    return SuccessResponse(200, authResp)
}
```

**Commit :**
```bash
git commit -m "feat: implémente le login endpoint

Code:
- auth_handler.go
- login route dans main.go

✅ Tests: 4/4 passent (GREEN)
✅ Couverture: 88%

Tests manuels:
✅ curl POST /api/auth/login (success → 200)
✅ curl POST /api/auth/login (wrong password → 401)
"
```

### Exemple Frontend : Formulaire de Login

#### 1. Tests (RED 🔴)

```tsx
// src/components/LoginForm.test.tsx
describe('LoginForm', () => {
  it('devrait afficher les champs email et password', () => {
    render(<LoginForm />);
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
  });

  it('devrait afficher une erreur si email invalide', async () => {
    render(<LoginForm />);
    const emailInput = screen.getByLabelText(/email/i);
    fireEvent.change(emailInput, { target: { value: 'invalid' } });
    fireEvent.submit(screen.getByRole('form'));

    await waitFor(() => {
      expect(screen.getByText(/email invalide/i)).toBeInTheDocument();
    });
  });

  it('devrait appeler onSubmit avec les bonnes données', async () => {
    const mockOnSubmit = jest.fn();
    render(<LoginForm onSubmit={mockOnSubmit} />);

    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: 'test@example.com' }
    });
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: 'password123' }
    });
    fireEvent.submit(screen.getByRole('form'));

    await waitFor(() => {
      expect(mockOnSubmit).toHaveBeenCalledWith({
        email: 'test@example.com',
        password: 'password123'
      });
    });
  });
});
```

**Commit :**
```bash
git commit -m "test: ajoute tests pour le LoginForm

Tests ajoutés:
- LoginForm.test.tsx (6 tests)
- Couvre: affichage, validation, soumission, erreurs

❌ Status: 0/6 tests passent (RED)
"
```

#### 2. Implémentation (GREEN 🟢)

```tsx
// src/components/LoginForm.tsx
export function LoginForm({ onSubmit }: LoginFormProps) {
  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: zodResolver(loginSchema)
  });

  return (
    <form role="form" onSubmit={handleSubmit(onSubmit)}>
      <label htmlFor="email">Email</label>
      <input {...register('email')} id="email" type="email" />
      {errors.email && <span>{errors.email.message}</span>}

      <label htmlFor="password">Mot de passe</label>
      <input {...register('password')} id="password" type="password" />
      {errors.password && <span>{errors.password.message}</span>}

      <button type="submit">Se connecter</button>
    </form>
  );
}
```

**Commit :**
```bash
git commit -m "feat: implémente le LoginForm

Code:
- LoginForm.tsx
- loginSchema validation avec Zod

✅ Tests: 6/6 passent (GREEN)
✅ Couverture: 100%

Tests manuels:
✅ Affichage correct dans le navigateur
✅ Validation fonctionne (email invalide → erreur)
"
```

## 🚫 Anti-patterns à éviter

### ❌ Écrire le code avant les tests
```bash
# MAUVAIS
git log
  feat: implémente le login  # Code avant tests ❌
  test: ajoute tests login   # Tests après ❌
```

### ✅ Correct : Tests avant code
```bash
# BON
git log
  feat: implémente le login  # Code après tests ✅
  test: ajoute tests login   # Tests d'abord ✅
```

### ❌ Commiter tests et code ensemble
```bash
# MAUVAIS
git commit -m "feat: login avec tests"  # Tout ensemble ❌
```

### ✅ Correct : Séparer les commits
```bash
# BON
git commit -m "test: ajoute tests login"  # D'abord les tests ✅
git commit -m "feat: implémente login"    # Puis le code ✅
```

### ❌ Tests qui passent dès le début
Si les tests passent immédiatement, c'est que :
- Le code existait déjà (pas de TDD)
- Les tests ne testent rien (assertions manquantes)
- Les tests sont mal écrits

### ✅ Correct : Tests échouent puis passent
```bash
# Étape 1 : RED
npm test
❌ 0/5 tests passent

git commit -m "test: ..."

# Étape 2 : GREEN
npm test
✅ 5/5 tests passent

git commit -m "feat: ..."
```

## 🎯 Checklist TDD

Avant chaque commit, vérifier :

**Commit de tests (test:) :**
- [ ] Les tests sont écrits
- [ ] Les tests ÉCHOUENT (RED)
- [ ] Tous les cas sont couverts (success + errors)
- [ ] Commit avec préfixe `test:`
- [ ] Message indique "❌ Status: Tests échouent (RED)"

**Commit de code (feat:/fix:) :**
- [ ] Le code est écrit
- [ ] Tous les tests PASSENT (GREEN)
- [ ] Couverture >= 80%
- [ ] Tests manuels effectués
- [ ] Commit avec préfixe `feat:` ou `fix:`
- [ ] Message indique "✅ Tests: X/X passent (GREEN)"

**Commit de refactoring (refactor:) - Optionnel :**
- [ ] Code amélioré
- [ ] Tous les tests PASSENT encore
- [ ] Commit avec préfixe `refactor:`
- [ ] Message indique "✅ Tests: X/X passent (toujours GREEN)"

## 📚 Ressources

- **TDD by Example** - Kent Beck
- **Growing Object-Oriented Software, Guided by Tests** - Steve Freeman & Nat Pryce
- **Red-Green-Refactor Cycle** - Martin Fowler

## 🤝 Questions fréquentes

### Q : Et si je ne sais pas comment écrire le test ?
**R :** C'est justement l'intérêt du TDD ! Si vous ne savez pas écrire le test, c'est que la spécification n'est pas claire. Clarifiez d'abord ce que vous voulez obtenir.

### Q : Dois-je vraiment commiter les tests qui échouent ?
**R :** **OUI !** C'est le principe du TDD. Cela garantit que :
1. Les tests existaient AVANT le code
2. L'historique Git est clair
3. On ne peut pas oublier d'écrire les tests

### Q : Que faire si j'ai déjà écrit du code sans tests ?
**R :** Deux options :
1. **Recommandé :** Supprimer le code, écrire les tests, ré-écrire le code
2. **Acceptable :** Écrire les tests maintenant, vérifier qu'ils échouent sans le code, puis commiter séparément

### Q : Combien de tests dans un commit ?
**R :** Dépend de la fonctionnalité. Généralement :
- Petite fonction : 3-5 tests
- Service complet : 10-20 tests
- Feature complète : Peut nécessiter plusieurs cycles TDD

---

**Dernière mise à jour :** 30/12/2025
**Version :** 1.0
**Statut :** MÉTHODOLOGIE OBLIGATOIRE ✅
