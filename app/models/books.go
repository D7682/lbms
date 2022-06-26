package models

// Book ..
type Book struct {
	ID     int64  `json:"id" bson:"_id"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
}

// BookRepository ..
type BookRepository interface {
	Save(book Book) error
	Get(id int64) (Book, error)
}
