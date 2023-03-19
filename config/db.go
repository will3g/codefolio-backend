package config

import (
	"context"

	"github.com/will3g/codefolio-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPortfolio(db *mongo.Database) ([]models.Portfolio, error) {
	collection := db.Collection("portfolio")

	var portfolio []models.Portfolio

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var p models.Portfolio
		if err := cur.Decode(&p); err != nil {
			return nil, err
		}

		portfolio = append(portfolio, p)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return portfolio, nil
}

func GetArticles(db *mongo.Database) ([]models.Article, error) {
	collection := db.Collection("articles")

	var articles []models.Article

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var a models.Article
		if err := cur.Decode(&a); err != nil {
			return nil, err
		}

		articles = append(articles, a)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
