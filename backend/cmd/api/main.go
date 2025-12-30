package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arnaud-dars/collec-app/internal/config"
	"github.com/arnaud-dars/collec-app/internal/handler"
	"github.com/arnaud-dars/collec-app/internal/middleware"
	"github.com/arnaud-dars/collec-app/internal/models"
	"github.com/arnaud-dars/collec-app/internal/repository"
	"github.com/arnaud-dars/collec-app/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const Version = "0.2.0"

func main() {
	fmt.Printf("Collec-App Backend v%s\n", Version)
	fmt.Println("Initializing server...")

	// Charger la configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Connexion à la base de données
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("✓ Database connected")

	// Auto-migration (pour le développement)
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	fmt.Println("✓ Migrations completed")

	// Initialiser les repositories
	userRepo := repository.NewUserRepository(db)

	// Initialiser les services
	authService := service.NewAuthService(
		userRepo,
		cfg.JWT.Secret,
		time.Duration(cfg.JWT.AccessTokenTTL)*time.Minute,
		time.Duration(cfg.JWT.RefreshTokenTTL)*time.Hour,
	)

	// Initialiser les handlers
	authHandler := handler.NewAuthHandler(authService)

	// Initialiser les middlewares
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Configurer les routes
	mux := http.NewServeMux()

	// Routes publiques
	mux.HandleFunc("/api/auth/register", authHandler.Register)
	mux.HandleFunc("/api/auth/login", authHandler.Login)
	mux.HandleFunc("/api/auth/refresh", authHandler.RefreshToken)
	mux.HandleFunc("/api/auth/logout", authHandler.Logout)

	// Routes protégées
	mux.HandleFunc("/api/auth/me", authMiddleware.RequireAuth(authHandler.GetMe))

	// Route de santé
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","version":"` + Version + `"}`))
	})

	// Démarrer le serveur
	addr := ":" + cfg.Server.Port
	fmt.Printf("✓ Server listening on http://localhost%s\n", addr)
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  POST   /api/auth/register")
	fmt.Println("  POST   /api/auth/login")
	fmt.Println("  POST   /api/auth/refresh")
	fmt.Println("  POST   /api/auth/logout")
	fmt.Println("  GET    /api/auth/me (protected)")
	fmt.Println("  GET    /health")

	if err := http.ListenAndServe(addr, enableCORS(mux)); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// initDatabase initialise la connexion à la base de données
func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// enableCORS ajoute les headers CORS pour le développement
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
