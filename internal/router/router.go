package router

import (
	"net/http"

	"echo-crate/internal/handlers"
	"echo-crate/internal/middleware"
	"echo-crate/internal/services"
	"echo-crate/internal/sessions"
	page "echo-crate/internal/views/pages"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		page.NotFound().Render(r.Context(), w)
		return
	}
	page.Index().Render(r.Context(), w)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	user, _ := sessions.GetAuthenticatedUser(r)
	page.Dashboard(user).Render(r.Context(), w)
}

func New(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.CleanPath)

	// Serve static files
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))
	r.Handle("/assets/*", fs)

	// Init Services
	authService := services.NewAuthService(db)
	todoService := services.NewTodoService(db)

	// Init Handlers
	authHandler := handlers.NewAuthHandler(authService)
	todoHandler := handlers.NewTodoHandler(todoService)

	// Routes
	r.Get("/", Index)

	// Auth routes
	r.Get("/auth", authHandler.Authenticate)
	r.Get("/auth/callback", authHandler.Callback)
	r.Get("/logout", authHandler.Logout)
	r.Get("/login", authHandler.LoginPage)

	// Todo routes
	r.Get("/todos", todoHandler.RenderPage)
	r.Post("/create-todo", todoHandler.Create)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/dashboard", Dashboard)
	})

	return r
}
