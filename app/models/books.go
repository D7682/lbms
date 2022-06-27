package models

import "context"

// Book ..
type Book struct {
	ID     int64  `json:"id" bson:"_id"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
}

// BookRepository ..
type BookRepository interface {
	Save(context.Context, Book) error
	Get(context.Context, int64) (Book, error)
	All(context.Context) ([]Book, error)
}
