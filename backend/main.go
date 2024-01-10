package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WieseChristoph/go-react-template/backend/internal/routes"
	"github.com/WieseChristoph/go-react-template/backend/internal/utils"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"golang.org/x/oauth2"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Get port from environment
	PORT, envExists := os.LookupEnv("PORT")
	if !envExists {
		PORT = "80"
	}

	// Connect to database
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))
	db, err := utils.ConnectDB(os.Getenv("DB_DRIVER"), connString)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate database
	err = utils.MigrateDB(db)
	if err != nil {
		log.Fatal(err)
	}

	// Setup Cronjobs
	c := cron.New()
	c.AddFunc("@hourly", utils.CleanExpiredSessions(db))
	c.Start()

	// Get app url from environment
	APP_URL, envExists := os.LookupEnv("APP_URL")
	if !envExists {
		APP_URL = "http://localhost"
	}

	// Create Discord OAuth2 config
	discordOauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("DISCORD_CLIENT_ID"),
		ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
		RedirectURL:  APP_URL + "/api/auth/discord/callback",
		Scopes:       []string{"identify", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}

	// Generate OAuth2 verifier
	oauthVerifier := oauth2.GenerateVerifier()

	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.AllowContentType("application/json", "text/xml"))
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "UPDATE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Routes
	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", routes.AuthResource{DB: db, OauthConfig: discordOauthConfig, Oauth2Verifier: oauthVerifier}.Routes())
		r.Mount("/users", routes.UsersResource{DB: db}.Routes())
		r.Mount("/todos", routes.TodosResource{DB: db}.Routes())
	})

	log.Printf("Server listening on port %s\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}
