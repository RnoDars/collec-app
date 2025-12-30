package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	appErrors "github.com/arnaud-dars/collec-app/internal/errors"
	"github.com/arnaud-dars/collec-app/internal/service"
)

// AuthMiddleware gère l'authentification JWT
type AuthMiddleware struct {
	authService service.AuthService
}

// NewAuthMiddleware crée une nouvelle instance de AuthMiddleware
func NewAuthMiddleware(authService service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// RequireAuth vérifie que l'utilisateur est authentifié
func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extraire le token du header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			m.respondWithError(w, http.StatusUnauthorized, appErrors.ErrUnauthorized.Code, "Token manquant")
			return
		}

		// Le format attendu est : "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			m.respondWithError(w, http.StatusUnauthorized, appErrors.ErrUnauthorized.Code, "Format de token invalide")
			return
		}

		token := parts[1]

		// Valider le token
		claims, err := m.authService.ValidateToken(token)
		if err != nil {
			m.respondWithError(w, http.StatusUnauthorized, appErrors.ErrUnauthorized.Code, "Token invalide ou expiré")
			return
		}

		// Ajouter l'user ID au contexte pour utilisation dans les handlers
		ctx := context.WithValue(r.Context(), "userID", claims.UserID.String())
		ctx = context.WithValue(ctx, "userEmail", claims.Email)

		// Passer à la suite avec le contexte enrichi
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// respondWithError envoie une réponse d'erreur JSON
func (m *AuthMiddleware) respondWithError(w http.ResponseWriter, status int, code, message string) {
	errorResponse := map[string]interface{}{
		"error": map[string]string{
			"code":    code,
			"message": message,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse)
}
