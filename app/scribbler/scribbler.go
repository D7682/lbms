package scribbler

import (
	"lbms/app/models"
	"os"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Scribbler struct {
	scribbler *scribble.Driver
}

func NewScribbler() (*Scribbler, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	db, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}

	return &Scribbler{
		scribbler: db,
	}, nil
}

func (s Scribbler) NewBook(book models.Book) error {
	err := s.scribbler.Write("books", book.Title, book)
	if err != nil {
		return err
	}
	return nil
}
