package routes

import (
	"database/sql"
	"net/http"

	"github.com/WieseChristoph/go-react-template/backend/internal/middleware"
	"github.com/WieseChristoph/go-react-template/backend/internal/models"
	"github.com/WieseChristoph/go-react-template/backend/internal/utils"

	"github.com/go-chi/chi/v5"
)

type UsersResource struct {
	DB *sql.DB
}

func (rs UsersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.AuthMiddleware(rs.DB))

	r.Get("/me", rs.Me)

	return r
}

func (rs UsersResource) Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)

	utils.SendResponse(w, http.StatusOK, true, "", user)
}
