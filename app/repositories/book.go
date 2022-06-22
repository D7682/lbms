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
	count, err := b.db.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	book.ID = count + 1

	_, err = b.db.InsertOne(ctx, book)
	return err
}

func (b BookRepo) Get(id int64) (models.Book, error) {
	ctx := context.Background()

	var book models.Book
	err := b.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&book)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}
