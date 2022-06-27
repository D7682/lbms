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

func (b BookRepo) Save(ctx context.Context, book models.Book) error {
	// https://stackoverflow.com/questions/61078884/mongodb-auto-increment-id-with-golang-mongo-driver
	count, err := b.db.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	book.ID = count + 1

	_, err = b.db.InsertOne(ctx, book)
	return err
}

func (b BookRepo) Get(ctx context.Context, id int64) (models.Book, error) {
	var book models.Book
	err := b.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&book)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b BookRepo) All(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	cursor, err := b.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
