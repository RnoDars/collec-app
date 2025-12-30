package errors

import (
	"fmt"
	"net/http"
)

// AppError représente une erreur applicative avec un code et un message
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
	Err        error  `json:"-"`
}

// Error implémente l'interface error
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s - %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap permet d'utiliser errors.Is et errors.As
func (e *AppError) Unwrap() error {
	return e.Err
}

// Erreurs communes

// NewAppError crée une nouvelle erreur applicative
func NewAppError(code, message string, statusCode int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

// Erreurs d'authentification
var (
	ErrUnauthorized = &AppError{
		Code:       "ERR_AUTH_001",
		Message:    "Non autorisé",
		StatusCode: http.StatusUnauthorized,
	}
	ErrInvalidCredentials = &AppError{
		Code:       "ERR_AUTH_002",
		Message:    "Identifiants invalides",
		StatusCode: http.StatusUnauthorized,
	}
	ErrTokenExpired = &AppError{
		Code:       "ERR_AUTH_003",
		Message:    "Token expiré",
		StatusCode: http.StatusUnauthorized,
	}
)

// Erreurs de validation
var (
	ErrValidation = &AppError{
		Code:       "ERR_VAL_001",
		Message:    "Erreur de validation",
		StatusCode: http.StatusBadRequest,
	}
	ErrInvalidInput = &AppError{
		Code:       "ERR_VAL_002",
		Message:    "Données d'entrée invalides",
		StatusCode: http.StatusBadRequest,
	}
)

// Erreurs de base de données
var (
	ErrDatabase = &AppError{
		Code:       "ERR_DB_001",
		Message:    "Erreur de base de données",
		StatusCode: http.StatusInternalServerError,
	}
	ErrNotFound = &AppError{
		Code:       "ERR_DB_002",
		Message:    "Ressource non trouvée",
		StatusCode: http.StatusNotFound,
	}
	ErrDuplicate = &AppError{
		Code:       "ERR_DB_003",
		Message:    "Ressource déjà existante",
		StatusCode: http.StatusConflict,
	}
)

// Erreurs génériques
var (
	ErrInternal = &AppError{
		Code:       "ERR_INT_001",
		Message:    "Erreur interne du serveur",
		StatusCode: http.StatusInternalServerError,
	}
	ErrForbidden = &AppError{
		Code:       "ERR_PERM_001",
		Message:    "Accès interdit",
		StatusCode: http.StatusForbidden,
	}
)

// WithError ajoute une erreur wrappée à une AppError
func (e *AppError) WithError(err error) *AppError {
	newErr := *e
	newErr.Err = err
	return &newErr
}

// WithMessage personnalise le message d'une AppError
func (e *AppError) WithMessage(message string) *AppError {
	newErr := *e
	newErr.Message = message
	return &newErr
}
