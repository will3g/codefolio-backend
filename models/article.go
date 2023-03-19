package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Path      string             `json:"path" bson:"path"`
	Section   string             `json:"section" bson:"section"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewArticle(title, content, path, section string) *Article {
	return &Article{
		ID:        primitive.NewObjectID(),
		Title:     title,
		Content:   content,
		Path:      path,
		Section:   section,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// func UpdateArticle(title, content, path, section string) *Article {
// 	return &Article{
// 		Title:     title,
// 		Content:   content,
// 		Path:      path,
// 		Section:   section,
// 		UpdatedAt: time.Now(),
// 	}
// }
