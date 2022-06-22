package models

func NewBook(title, author string) *Book {
	return &Book{
		Title:  title,
		Author: author,
	}
}
