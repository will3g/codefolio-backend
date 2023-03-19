package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/will3g/codefolio-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetArticles(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, err := db.Collection("articles").Find(context.Background(), bson.D{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.Background())

		var results []models.Article
		for cursor.Next(context.Background()) {
			var article models.Article
			err := cursor.Decode(&article)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, article)
		}

		if err := cursor.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetArticle(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var getArticle struct {
			Path string `json:"path"`
		}

		if err := json.NewDecoder(r.Body).Decode(&getArticle); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		cursor, err := db.Collection("articles").Find(context.Background(), bson.D{{Key: "path", Value: getArticle.Path}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.Background())

		var results []models.Article
		for cursor.Next(context.Background()) {
			var article models.Article
			err := cursor.Decode(&article)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, article)
		}

		if err := cursor.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CreateArticle(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newArticle struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			Path    string `json:"path"`
			Section string `json:"section"`
		}

		if err := json.NewDecoder(r.Body).Decode(&newArticle); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		article := models.NewArticle(
			newArticle.Title,
			newArticle.Content,
			newArticle.Path,
			newArticle.Section)

		collection := db.Collection("articles")
		if _, err := collection.InsertOne(context.Background(), article); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error creating article"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(article)
	}
}

func UpdateArticle(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateArticle struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			Path    string `json:"path"`
			Section string `json:"section"`
		}

		if err := json.NewDecoder(r.Body).Decode(&updateArticle); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		bsonStructure := bson.D{
			{Key: "title", Value: updateArticle.Title},
			{Key: "content", Value: updateArticle.Content},
			{Key: "path", Value: updateArticle.Path},
			{Key: "section", Value: updateArticle.Section},
			{Key: "updated_at", Value: time.Now()},
		}

		filter := bson.D{{Key: "path", Value: updateArticle.Path}}
		update := bson.M{"$set": bsonStructure}

		collection := db.Collection("articles")
		if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error updating article"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

func DeleteArticle(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deleteArticle struct {
			Path string `json:"path"`
		}

		if err := json.NewDecoder(r.Body).Decode(&deleteArticle); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		filter := bson.D{{Key: "path", Value: deleteArticle.Path}}

		collection := db.Collection("articles")
		if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error deletion article"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

// func DeleteArticle(db *mongo.Database) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		id := vars["id"]

// 		err := db.DeleteArticle(db, id)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("Error deleting article"))
// 			return
// 		}

// 		w.Write([]byte("Article deleted successfully"))
// 	}
// }
