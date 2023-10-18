package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ptflp/go-test/app/handler"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func PrepareRoutes(handleTodo handler.TodosHandler) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api", func(r chi.Router) {
		r.Route("/todos", func(r chi.Router) {
			todos := handleTodo
			r.Post("/create", todos.CreateTodo)
			r.Post("/complete/{id}", todos.CompleteTodo)
			r.Get("/all", todos.GetAll)
			r.Post("/delete/{id}", todos.Delete)
			r.Post("/update", todos.Update)
		})
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)

	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
