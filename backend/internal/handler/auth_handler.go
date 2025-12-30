package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arnaud-dars/collec-app/internal/dto"
	appErrors "github.com/arnaud-dars/collec-app/internal/errors"
	"github.com/arnaud-dars/collec-app/internal/service"
	"github.com/go-playground/validator/v10"
)

// AuthHandler gère les endpoints d'authentification
type AuthHandler struct {
	authService service.AuthService
	validate    *validator.Validate
}

// NewAuthHandler crée une nouvelle instance de AuthHandler
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

// Register gère l'inscription d'un nouvel utilisateur
// POST /api/auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	// Décoder le body JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, appErrors.ErrValidation.Code, "Données invalides", err)
		return
	}

	// Valider les données
	if err := h.validate.Struct(req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, appErrors.ErrValidation.Code, "Erreur de validation", err)
		return
	}

	// Créer l'utilisateur
	user, err := h.authService.Register(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			h.respondWithError(w, http.StatusConflict, "ERR_AUTH_002", "Cet email est déjà utilisé", err)
			return
		}
		if errors.Is(err, service.ErrWeakPassword) {
			h.respondWithError(w, http.StatusBadRequest, "ERR_AUTH_003", err.Error(), err)
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "ERR_INTERNAL_001", "Erreur lors de la création du compte", err)
		return
	}

	// Générer les tokens pour auto-login après inscription
	accessToken, refreshToken, _, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "ERR_INTERNAL_001", "Compte créé mais erreur de connexion", err)
		return
	}

	// Réponse
	response := dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         dto.ToUserDTO(user),
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

// Login gère la connexion d'un utilisateur
// POST /api/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	// Décoder le body JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, appErrors.ErrValidation.Code, "Données invalides", err)
		return
	}

	// Valider les données
	if err := h.validate.Struct(req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, appErrors.ErrValidation.Code, "Erreur de validation", err)
		return
	}

	// Authentifier l'utilisateur
	accessToken, refreshToken, user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			h.respondWithError(w, http.StatusUnauthorized, appErrors.ErrUnauthorized.Code, "Email ou mot de passe incorrect", err)
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "ERR_INTERNAL_001", "Erreur lors de la connexion", err)
		return
	}

	// Réponse
	response := dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         dto.ToUserDTO(user),
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// RefreshToken génère un nouveau access token
// POST /api/auth/refresh
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req dto.RefreshTokenRequest

	// Décoder le body JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, appErrors.ErrValidation.Code, "Données invalides", err)
		return
	}

	// Valider les données
	if err := h.validate.Struct(req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, appErrors.ErrValidation.Code, "Erreur de validation", err)
		return
	}

	// Générer un nouveau access token
	accessToken, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, service.ErrInvalidToken) {
			h.respondWithError(w, http.StatusUnauthorized, appErrors.ErrUnauthorized.Code, "Token invalide ou expiré", err)
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "ERR_INTERNAL_001", "Erreur lors du refresh du token", err)
		return
	}

	// Réponse
	response := map[string]string{
		"accessToken": accessToken,
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// GetMe retourne les informations de l'utilisateur connecté
// GET /api/auth/me (route protégée)
func (h *AuthHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	// L'user ID est injecté dans le contexte par le middleware d'authentification
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		h.respondWithError(w, http.StatusUnauthorized, appErrors.ErrUnauthorized.Code, "Non authentifié", nil)
		return
	}

	// TODO: Récupérer l'utilisateur depuis le repository
	// Pour l'instant, on retourne juste l'ID
	response := map[string]string{
		"userId": userID,
		"message": "Endpoint GetMe - À implémenter avec le repository",
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// Logout déconnecte l'utilisateur (côté client)
// POST /api/auth/logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// La déconnexion est gérée côté client (suppression des tokens)
	// Ici on peut juste confirmer la déconnexion
	response := map[string]string{
		"message": "Déconnexion réussie",
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// respondWithJSON envoie une réponse JSON
func (h *AuthHandler) respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// respondWithError envoie une réponse d'erreur JSON
func (h *AuthHandler) respondWithError(w http.ResponseWriter, status int, code, message string, err error) {
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
