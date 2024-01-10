package middleware

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/WieseChristoph/go-react-template/backend/internal/models"
	"github.com/WieseChristoph/go-react-template/backend/internal/utils"
)

func AuthMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get session token from cookie
			sessionTokenCookie, err := r.Cookie("session_token")
			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized, false, "Session token missing", nil)
				return
			}
			sessionToken := sessionTokenCookie.Value

			// Get session from database
			session, err := models.GetSessionByToken(db, sessionToken)
			if err != nil {
				utils.SendResponse(w, http.StatusUnauthorized, false, "Invalid session token", nil)
				return
			}

			// Check if session is expired
			if session.IsExpired() {
				session.Delete(db)
				utils.SendResponse(w, http.StatusUnauthorized, false, "Session expired", nil)
				return
			}

			// Add session to context
			ctx := context.WithValue(r.Context(), "session", session)

			// Get user from database
			user, err := models.GetUserById(db, session.UserId)
			if err != nil {
				utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
				return
			}

			// Add user to context
			ctx = context.WithValue(ctx, "user", user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
