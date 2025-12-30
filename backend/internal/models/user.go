package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User représente un utilisateur de l'application
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"` // Le tag json:"-" empêche l'export du password en JSON
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BeforeCreate hook GORM pour générer un UUID avant la création
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// TableName spécifie le nom de la table en base de données
func (User) TableName() string {
	return "users"
}
