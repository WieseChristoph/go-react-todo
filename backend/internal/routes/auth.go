package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/WieseChristoph/go-react-template/backend/internal/middleware"
	"github.com/WieseChristoph/go-react-template/backend/internal/models"
	"github.com/WieseChristoph/go-react-template/backend/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/oauth2"

	"github.com/go-chi/chi/v5"
)

const TOKEN_EXPIRATION = 7 * 24 * time.Hour

type AuthResource struct {
	DB             *sql.DB
	OauthConfig    *oauth2.Config
	Oauth2Verifier string
}

func (rs AuthResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(rs.DB))
		r.Post("/logout", rs.Logout)
	})

	r.Route("/discord", func(r chi.Router) {
		r.Get("/login", rs.Login)
		r.Get("/callback", rs.Callback)
	})

	return r
}

func (rs AuthResource) Login(w http.ResponseWriter, r *http.Request) {
	// Redirect user to Discord OAuth2 login page
	// TODO: Add state to prevent CSRF
	url := rs.OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(rs.Oauth2Verifier))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (rs AuthResource) Callback(w http.ResponseWriter, r *http.Request) {
	// Exchange code for token
	token, err := rs.OauthConfig.Exchange(context.Background(), r.FormValue("code"), oauth2.VerifierOption(rs.Oauth2Verifier))
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	// Get user from Discord API
	res, err := rs.OauthConfig.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	defer res.Body.Close()

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Parse user from response body
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Save user to database
	user, err = user.Save(rs.DB)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Create session
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(TOKEN_EXPIRATION)
	session := models.Session{
		Token:     sessionToken,
		UserId:    user.ID,
		ExpiresAt: expiresAt,
	}

	// Save session to database
	session, err = session.Save(rs.DB)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Set session token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Expires:  session.ExpiresAt,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (rs AuthResource) Logout(w http.ResponseWriter, r *http.Request) {
	// Get session from context
	session := r.Context().Value("session").(models.Session)

	// Delete session from database
	err := session.Delete(rs.DB)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Delete session token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now(),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	})

	utils.SendResponse(w, http.StatusOK, true, "", nil)
}
