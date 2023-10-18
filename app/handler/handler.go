package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ptflp/go-test/app/entity"
	"github.com/ptflp/go-test/app/service"
	"net/http"
	"strconv"
)

type TodosHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
	CompleteTodo(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Todos struct {
	service service.Todoer
}

func NewTodosHandler(serviceTodo service.Todoer) TodosHandler {
	return &Todos{
		service: serviceTodo,
	}
}

func (h *Todos) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// parse json from request body
	var todo entity.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.Create(r.Context(), &todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CompleteTodo marks as complete
func (h *Todos) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	// parse {id: 1} from go chi router
	id := extractID(r)
	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.service.Complete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetAll get all todos
func (h *Todos) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(todos)
}

// Delete - deletes todo
func (h *Todos) Delete(w http.ResponseWriter, r *http.Request) {
	id := extractID(r)
	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.service.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Update todo
func (h *Todos) Update(w http.ResponseWriter, r *http.Request) {
	// parse json from request body
	var todo entity.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.Update(r.Context(), &todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func extractID(r *http.Request) int {
	rawID := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(rawID)

	return id
}
