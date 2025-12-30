package service

import (
	"errors"
	"testing"
	"time"

	"github.com/arnaud-dars/collec-app/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mock du UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id uuid.UUID) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) ExistsByEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

// Helper function pour générer un token de test
func generateTestToken(authSvc AuthService, userID uuid.UUID, email string, duration time.Duration) (string, error) {
	svc, ok := authSvc.(*authService)
	if !ok {
		return "", errors.New("invalid auth service type")
	}
	return svc.generateToken(userID, email, duration)
}

// Tests du service Auth

func TestRegister_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	email := "test@example.com"
	password := "password123"

	mockRepo.On("ExistsByEmail", email).Return(false, nil)
	mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	// Act
	user, err := authService.Register(email, password)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, user.Password)
	mockRepo.AssertExpectations(t)
}

func TestRegister_EmailAlreadyExists(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	email := "existing@example.com"
	password := "password123"

	mockRepo.On("ExistsByEmail", email).Return(true, nil)

	// Act
	user, err := authService.Register(email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrEmailAlreadyExists, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestRegister_WeakPassword(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	email := "test@example.com"
	password := "weak" // Moins de 8 caractères

	// Act
	user, err := authService.Register(email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrWeakPassword, err)
	assert.Nil(t, user)
}

func TestLogin_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	email := "test@example.com"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	existingUser := &models.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
	}

	mockRepo.On("FindByEmail", email).Return(existingUser, nil)

	// Act
	accessToken, refreshToken, user, err := authService.Login(email, password)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidCredentials_UserNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	email := "nonexistent@example.com"
	password := "password123"

	mockRepo.On("FindByEmail", email).Return(nil, nil)

	// Act
	accessToken, refreshToken, user, err := authService.Login(email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Empty(t, accessToken)
	assert.Empty(t, refreshToken)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidCredentials_WrongPassword(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	email := "test@example.com"
	password := "password123"
	wrongPassword := "wrongpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	existingUser := &models.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
	}

	mockRepo.On("FindByEmail", email).Return(existingUser, nil)

	// Act
	accessToken, refreshToken, user, err := authService.Login(email, wrongPassword)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Empty(t, accessToken)
	assert.Empty(t, refreshToken)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestValidateToken_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	userID := uuid.New()
	email := "test@example.com"

	// Générer un token valide
	token, err := generateTestToken(authService, userID, email, 15*time.Minute)
	assert.NoError(t, err)

	// Act
	claims, err := authService.ValidateToken(token)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
}

func TestValidateToken_InvalidToken(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	invalidToken := "invalid.token.here"

	// Act
	claims, err := authService.ValidateToken(invalidToken)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidToken, err)
	assert.Nil(t, claims)
}

func TestRefreshToken_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	userID := uuid.New()
	email := "test@example.com"

	existingUser := &models.User{
		ID:    userID,
		Email: email,
	}

	// Générer un refresh token valide
	refreshToken, err := generateTestToken(authService, userID, email, 168*time.Hour)
	assert.NoError(t, err)

	mockRepo.On("FindByID", userID).Return(existingUser, nil)

	// Act
	newAccessToken, err := authService.RefreshToken(refreshToken)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, newAccessToken)
	mockRepo.AssertExpectations(t)
}

func TestRefreshToken_UserNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	userID := uuid.New()
	email := "test@example.com"

	// Générer un refresh token valide
	refreshToken, err := generateTestToken(authService, userID, email, 168*time.Hour)
	assert.NoError(t, err)

	mockRepo.On("FindByID", userID).Return(nil, nil)

	// Act
	newAccessToken, err := authService.RefreshToken(refreshToken)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidToken, err)
	assert.Empty(t, newAccessToken)
	mockRepo.AssertExpectations(t)
}

func TestRegister_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	authService := NewAuthService(mockRepo, "test-secret-key", 15*time.Minute, 168*time.Hour)

	email := "test@example.com"
	password := "password123"
	dbError := errors.New("database error")

	mockRepo.On("ExistsByEmail", email).Return(false, dbError)

	// Act
	user, err := authService.Register(email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, dbError, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}
