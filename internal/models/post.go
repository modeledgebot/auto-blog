package models

import (
    "time"
)

// Category represents a main section of the blog
type Category struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Posts       []Post    `json:"posts,omitempty"`
}

// Tag represents a way to group related posts
type Tag struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Count int    `json:"count"` // Number of posts using this tag
}

// Post represents a blog post
type Post struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Summary     string    `json:"summary"`
    Content     string    `json:"content"`
    Category    Category  `json:"category"`     // Each post belongs to one main category
    Tags        []Tag     `json:"tags"`         // Posts can have multiple tags
    RelatedPosts []string `json:"related_posts"` // IDs of related posts (based on tags)
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// Navigation represents the blog's navigation structure
type Navigation struct {
    Categories []Category          `json:"categories"`
    TagCloud   map[string]int     `json:"tag_cloud"`    // Tag name -> post count
    Recent     []Post             `json:"recent_posts"` // Recent posts across all categories
}