package repository

import (
	"errors"

	"github.com/arnaud-dars/collec-app/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository définit l'interface pour les opérations sur les utilisateurs
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uuid.UUID) (*models.User, error)
	ExistsByEmail(email string) (bool, error)
}

// userRepository implémente UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository crée une nouvelle instance de UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create insère un nouvel utilisateur en base de données
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByEmail recherche un utilisateur par son email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Pas d'erreur si non trouvé, juste nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByID recherche un utilisateur par son ID
func (r *userRepository) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Pas d'erreur si non trouvé, juste nil
		}
		return nil, err
	}
	return &user, nil
}

// ExistsByEmail vérifie si un email existe déjà en base de données
func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
