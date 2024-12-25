package main

import (
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/auto-blog/internal/handlers"
)

func main() {
    r := chi.NewRouter()
    
    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    
    // Serve static files
    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))
    
    // Blog routes
    r.Get("/", handlers.HomeHandler)
    r.Get("/post/{id}", handlers.PostHandler)
    r.Get("/category/{category}", handlers.CategoryHandler)
    r.Get("/tag/{tag}", handlers.TagHandler)
    
    // Admin routes
    r.Get("/admin", handlers.AdminHandler)
    r.Route("/admin/post/new", func(r chi.Router) {
        r.Get("/", handlers.NewPostHandler)
        r.Post("/", handlers.NewPostHandler)
    })
    r.Route("/admin/categories", func(r chi.Router) {
        r.Get("/", handlers.CategoriesHandler)
        r.Post("/", handlers.CategoriesHandler)
    })
    r.Route("/admin/tags", func(r chi.Router) {
        r.Get("/", handlers.TagsHandler)
        r.Post("/", handlers.TagsHandler)
    })
    
    // Theme toggle
    r.Post("/toggle-theme", handlers.ToggleThemeHandler)
    
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}