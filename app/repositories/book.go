package repositories

import (
	"context"
	"lbms/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepo struct {
	db *mongo.Collection
}

func NewBookRepo(db *mongo.Collection) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (b BookRepo) Save(book models.Book) error {
	// https://stackoverflow.com/questions/61078884/mongodb-auto-increment-id-with-golang-mongo-driver
	ctx := context.Background()
	_, err := b.db.InsertOne(ctx, book)
	return err
}

func (b BookRepo) GetAll() ([]models.Book, error) {
	ctx := context.Background()
	result, err := b.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var books []models.Book
	err = result.All(ctx, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
