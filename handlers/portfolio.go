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

func GetPortfolios(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, err := db.Collection("portfolio").Find(context.Background(), bson.D{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.Background())

		var results []models.Portfolio
		for cursor.Next(context.Background()) {
			var portfolio models.Portfolio
			err := cursor.Decode(&portfolio)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, portfolio)
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

func GetPortfolio(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var getArticle struct {
			Path string `json:"path"`
		}

		if err := json.NewDecoder(r.Body).Decode(&getArticle); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		cursor, err := db.Collection("portfolio").Find(context.Background(), bson.D{{Key: "path", Value: getArticle.Path}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.Background())

		var results []models.Portfolio
		for cursor.Next(context.Background()) {
			var portfolio models.Portfolio
			err := cursor.Decode(&portfolio)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, portfolio)
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

func CreatePortfolio(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPortfolio struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			Path        string   `json:"path"`
			Category    string   `json:"category"`
			Frameworks  []string `json:"frameworks" bson:"frameworks"`
			Images      []string `json:"images" bson:"images"`
		}

		if err := json.NewDecoder(r.Body).Decode(&newPortfolio); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		portfolio := models.NewPortfolio(
			newPortfolio.Title,
			newPortfolio.Description,
			newPortfolio.Path,
			newPortfolio.Category,
			newPortfolio.Frameworks,
			newPortfolio.Images)

		collection := db.Collection("portfolio")
		if _, err := collection.InsertOne(context.Background(), portfolio); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error creating portfolio"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(portfolio)
	}
}

func UpdatePortfolio(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatePortfolio struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			Path        string   `json:"path"`
			Category    string   `json:"category"`
			Frameworks  []string `json:"frameworks" bson:"frameworks"`
			Images      []string `json:"images" bson:"images"`
		}

		if err := json.NewDecoder(r.Body).Decode(&updatePortfolio); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		bsonStructure := bson.D{
			{Key: "title", Value: updatePortfolio.Title},
			{Key: "description", Value: updatePortfolio.Description},
			{Key: "path", Value: updatePortfolio.Path},
			{Key: "category", Value: updatePortfolio.Category},
			{Key: "frameworks", Value: updatePortfolio.Frameworks},
			{Key: "images", Value: updatePortfolio.Images},
			{Key: "updated_at", Value: time.Now()},
		}

		filter := bson.D{{Key: "path", Value: updatePortfolio.Path}}
		update := bson.M{"$set": bsonStructure}

		collection := db.Collection("portfolio")
		if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error updating portfolio"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

func DeletePortfolio(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deletePortfolio struct {
			Path string `json:"path"`
		}

		if err := json.NewDecoder(r.Body).Decode(&deletePortfolio); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request payload"))
			return
		}

		filter := bson.D{{Key: "path", Value: deletePortfolio.Path}}

		collection := db.Collection("portfolio")
		if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error deletion portfolio"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	}
}
