package dto

import (
	"time"

	"github.com/arnaud-dars/collec-app/internal/models"
	"github.com/google/uuid"
)

// RegisterRequest représente les données d'inscription
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// LoginRequest représente les données de connexion
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RefreshTokenRequest représente la demande de refresh token
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// AuthResponse représente la réponse après inscription ou connexion réussie
type AuthResponse struct {
	AccessToken  string  `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
	User         UserDTO `json:"user"`
}

// UserDTO représente les données publiques d'un utilisateur
type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// ToUserDTO convertit un modèle User en UserDTO
func ToUserDTO(user *models.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
