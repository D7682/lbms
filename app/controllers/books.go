package controllers

import (
	"fmt"
	"lbms/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BookHandler will hold the repo that deals with
// the books database
type BookHandler struct {
	bookRepo models.BookRepository
}

// NewBookHandler used to define a variable that will contain
// all the handlers that were created.
func NewBookHandler(bookrepo models.BookRepository) *BookHandler {
	return &BookHandler{
		bookRepo: bookrepo,
	}
}

// BookHandler.NewBook is the handler used for creating a new book
// in the book repo, used in the POST route
func (b BookHandler) NewBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err := b.bookRepo.Save(c, book)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

// BookHandler.Get is the handler used to retrieve a book by the id.
func (b BookHandler) Get(c *gin.Context) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	fmt.Println(idstr)

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	books, err := b.bookRepo.Get(c, int64(id))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, books)
}

// BookHandler.All retrieves all books from the book repository.
func (b BookHandler) All(c *gin.Context) {
	books, err := b.bookRepo.All(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, books)
}
