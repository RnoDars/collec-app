package service

import (
	"errors"
	"time"

	"github.com/arnaud-dars/collec-app/internal/models"
	"github.com/arnaud-dars/collec-app/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists = errors.New("cet email est déjà utilisé")
	ErrInvalidCredentials = errors.New("email ou mot de passe incorrect")
	ErrInvalidToken       = errors.New("token invalide")
	ErrWeakPassword       = errors.New("le mot de passe doit contenir au moins 8 caractères")
)

// JWTClaims représente les données contenues dans le JWT
type JWTClaims struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

// AuthService définit l'interface pour les opérations d'authentification
type AuthService interface {
	Register(email, password string) (*models.User, error)
	Login(email, password string) (accessToken, refreshToken string, user *models.User, err error)
	RefreshToken(refreshToken string) (string, error)
	ValidateToken(token string) (*JWTClaims, error)
}

// authService implémente AuthService
type authService struct {
	userRepo             repository.UserRepository
	jwtSecret            []byte
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

// NewAuthService crée une nouvelle instance de AuthService
func NewAuthService(
	userRepo repository.UserRepository,
	jwtSecret string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) AuthService {
	return &authService{
		userRepo:             userRepo,
		jwtSecret:            []byte(jwtSecret),
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

// Register crée un nouveau compte utilisateur
func (s *authService) Register(email, password string) (*models.User, error) {
	// Valider le mot de passe
	if len(password) < 8 {
		return nil, ErrWeakPassword
	}

	// Vérifier que l'email n'existe pas déjà
	exists, err := s.userRepo.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Créer l'utilisateur
	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login authentifie un utilisateur et retourne les tokens JWT
func (s *authService) Login(email, password string) (string, string, *models.User, error) {
	// Trouver l'utilisateur par email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", "", nil, err
	}
	if user == nil {
		return "", "", nil, ErrInvalidCredentials
	}

	// Vérifier le mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", nil, ErrInvalidCredentials
	}

	// Générer les tokens
	accessToken, err := s.generateToken(user.ID, user.Email, s.accessTokenDuration)
	if err != nil {
		return "", "", nil, err
	}

	refreshToken, err := s.generateToken(user.ID, user.Email, s.refreshTokenDuration)
	if err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}

// RefreshToken génère un nouveau access token à partir d'un refresh token valide
func (s *authService) RefreshToken(refreshToken string) (string, error) {
	// Valider le refresh token
	claims, err := s.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	// Vérifier que l'utilisateur existe toujours
	user, err := s.userRepo.FindByID(claims.UserID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrInvalidToken
	}

	// Générer un nouveau access token
	accessToken, err := s.generateToken(user.ID, user.Email, s.accessTokenDuration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// ValidateToken valide et parse un JWT
func (s *authService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Vérifier la méthode de signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// generateToken génère un JWT avec les claims spécifiés
func (s *authService) generateToken(userID uuid.UUID, email string, duration time.Duration) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
