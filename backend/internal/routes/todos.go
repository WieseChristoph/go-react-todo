package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/WieseChristoph/go-react-template/backend/internal/middleware"
	"github.com/WieseChristoph/go-react-template/backend/internal/models"
	"github.com/WieseChristoph/go-react-template/backend/internal/utils"

	"github.com/go-chi/chi/v5"
)

type TodosResource struct {
	DB *sql.DB
}

func (rs TodosResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.AuthMiddleware(rs.DB))

	r.Get("/", rs.List)
	r.Post("/", rs.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Post("/", rs.Update)
		r.Delete("/", rs.Delete)
	})

	return r
}

func (rs TodosResource) List(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	todos, err := models.GetAllTodosByUser(rs.DB, user.ID)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utils.SendResponse(w, http.StatusOK, true, "", todos)
}

func (rs TodosResource) Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	user := r.Context().Value("user").(models.User)
	todo.UserId = user.ID

	todo, err = todo.Save(rs.DB)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utils.SendResponse(w, http.StatusCreated, true, "", todo)
}

func (rs TodosResource) Get(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, "Invalid todo ID", nil)
		return
	}

	todo, err := models.GetTodoById(rs.DB, todoId)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Check if user is allowed to access this todo
	user := r.Context().Value("user").(models.User)
	if todo.UserId != user.ID {
		utils.SendResponse(w, http.StatusForbidden, false, "You are not allowed to access this todo", nil)
		return
	}

	utils.SendResponse(w, http.StatusOK, true, "", todo)
}

func (rs TodosResource) Update(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, "Invalid todo ID", nil)
		return
	}

	todo, err := models.GetTodoById(rs.DB, todoId)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Check if user is allowed to access this todo
	user := r.Context().Value("user").(models.User)
	if todo.UserId != user.ID {
		utils.SendResponse(w, http.StatusForbidden, false, "You are not allowed to access this todo", nil)
		return
	}

	var newTodo models.Todo
	err = json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	todo.Title = newTodo.Title
	todo.Description = newTodo.Description
	todo.Status = newTodo.Status

	todo, err = todo.Update(rs.DB)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utils.SendResponse(w, http.StatusCreated, true, "", todo)
}

func (rs TodosResource) Delete(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, "Invalid todo ID", nil)
		return
	}

	todo, err := models.GetTodoById(rs.DB, todoId)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// Check if user is allowed to access this todo
	user := r.Context().Value("user").(models.User)
	if todo.UserId != user.ID {
		utils.SendResponse(w, http.StatusForbidden, false, "You are not allowed to access this todo", nil)
		return
	}

	err = models.DeleteTodoById(rs.DB, todoId)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utils.SendResponse(w, http.StatusOK, true, "", nil)
}
