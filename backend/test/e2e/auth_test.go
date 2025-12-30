package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	baseURL = "http://localhost:8080"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"user"`
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func TestE2E_AuthenticationFlow(t *testing.T) {
	// Générer un email unique pour éviter les conflits
	timestamp := time.Now().Unix()
	testEmail := fmt.Sprintf("e2e-test-%d@example.com", timestamp)
	testPassword := "password123"

	t.Run("1. Inscription d'un nouvel utilisateur - devrait retourner 201", func(t *testing.T) {
		reqBody := RegisterRequest{
			Email:    testEmail,
			Password: testPassword,
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/register", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var authResp AuthResponse
		err = json.NewDecoder(resp.Body).Decode(&authResp)
		require.NoError(t, err)

		// Vérifier que les tokens sont présents
		assert.NotEmpty(t, authResp.AccessToken)
		assert.NotEmpty(t, authResp.RefreshToken)
		assert.Equal(t, testEmail, authResp.User.Email)
		assert.NotEmpty(t, authResp.User.ID)
		assert.False(t, authResp.User.CreatedAt.IsZero())
	})

	t.Run("2. Inscription avec email existant - devrait retourner 409", func(t *testing.T) {
		reqBody := RegisterRequest{
			Email:    testEmail,
			Password: testPassword,
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/register", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusConflict, resp.StatusCode)

		var errResp ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		require.NoError(t, err)

		assert.Equal(t, "ERR_AUTH_002", errResp.Error.Code)
		assert.Contains(t, errResp.Error.Message, "Cet email est déjà utilisé")
	})

	t.Run("3. Connexion avec identifiants valides - devrait retourner 200", func(t *testing.T) {
		reqBody := LoginRequest{
			Email:    testEmail,
			Password: testPassword,
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/login", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var authResp AuthResponse
		err = json.NewDecoder(resp.Body).Decode(&authResp)
		require.NoError(t, err)

		assert.NotEmpty(t, authResp.AccessToken)
		assert.NotEmpty(t, authResp.RefreshToken)
		assert.Equal(t, testEmail, authResp.User.Email)
	})

	t.Run("4. Connexion avec mauvais mot de passe - devrait retourner 401", func(t *testing.T) {
		reqBody := LoginRequest{
			Email:    testEmail,
			Password: "wrongpassword",
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/login", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		var errResp ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		require.NoError(t, err)

		assert.Equal(t, "ERR_AUTH_001", errResp.Error.Code)
		assert.Contains(t, errResp.Error.Message, "Email ou mot de passe incorrect")
	})

	t.Run("5. Connexion avec email inexistant - devrait retourner 401", func(t *testing.T) {
		reqBody := LoginRequest{
			Email:    "nonexistent@example.com",
			Password: testPassword,
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/login", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	var validAccessToken string

	t.Run("6. Récupérer un token valide pour les tests suivants", func(t *testing.T) {
		reqBody := LoginRequest{
			Email:    testEmail,
			Password: testPassword,
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/login", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		var authResp AuthResponse
		err = json.NewDecoder(resp.Body).Decode(&authResp)
		require.NoError(t, err)

		validAccessToken = authResp.AccessToken
		require.NotEmpty(t, validAccessToken)
	})

	t.Run("7. Accès à endpoint protégé avec token valide - devrait retourner 200", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/api/auth/me", nil)
		require.NoError(t, err)

		req.Header.Set("Authorization", "Bearer "+validAccessToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Le endpoint retourne actuellement un message temporaire
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)

		assert.NotEmpty(t, result["userId"])
	})

	t.Run("8. Accès à endpoint protégé sans token - devrait retourner 401", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/api/auth/me", nil)
		require.NoError(t, err)

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("9. Accès à endpoint protégé avec token invalide - devrait retourner 401", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/api/auth/me", nil)
		require.NoError(t, err)

		req.Header.Set("Authorization", "Bearer invalid-token")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

func TestE2E_ValidationErrors(t *testing.T) {
	t.Run("Inscription sans email - devrait retourner 400", func(t *testing.T) {
		reqBody := map[string]string{
			"password": "password123",
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/register", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Inscription avec mot de passe trop court - devrait retourner 400", func(t *testing.T) {
		reqBody := RegisterRequest{
			Email:    "test@example.com",
			Password: "short",
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/register", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Inscription avec email invalide - devrait retourner 400", func(t *testing.T) {
		reqBody := RegisterRequest{
			Email:    "invalid-email",
			Password: "password123",
		}

		body, _ := json.Marshal(reqBody)
		resp, err := http.Post(baseURL+"/api/auth/register", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestE2E_HealthCheck(t *testing.T) {
	t.Run("Endpoint /health - devrait retourner 200", func(t *testing.T) {
		resp, err := http.Get(baseURL + "/health")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)

		assert.Equal(t, "ok", result["status"])
	})
}
