package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/auto-blog/internal/handlers"
)

func main() {
    r := mux.NewRouter()
    
    // Serve static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))
    
    // Blog routes
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/post/{id}", handlers.PostHandler).Methods("GET")
    r.HandleFunc("/category/{category}", handlers.CategoryHandler).Methods("GET")
    r.HandleFunc("/tag/{tag}", handlers.TagHandler).Methods("GET")
    
    // Admin routes
    r.HandleFunc("/admin", handlers.AdminHandler).Methods("GET")
    r.HandleFunc("/admin/post/new", handlers.NewPostHandler).Methods("GET", "POST")
    r.HandleFunc("/admin/categories", handlers.CategoriesHandler).Methods("GET", "POST")
    r.HandleFunc("/admin/tags", handlers.TagsHandler).Methods("GET", "POST")
    
    // Theme toggle
    r.HandleFunc("/toggle-theme", handlers.ToggleThemeHandler).Methods("POST")
    
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}