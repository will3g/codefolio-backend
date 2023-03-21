package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Portfolio struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Path        string             `json:"path" bson:"path"`
	Category    string             `json:"category" bson:"category"`
	Frameworks  []string           `json:"frameworks" bson:"frameworks"`
	Images      []string           `json:"images" bson:"images"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewPortfolio(
	title, description, path, category string,
	frameworks, images []string) *Portfolio {
	return &Portfolio{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Description: description,
		Path:        path,
		Category:    category,
		Frameworks:  frameworks,
		Images:      images,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
