package controllers

import (
	"lbms/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookRepo models.BookRepository
}

func NewBookHandler(bookrepo models.BookRepository) *BookHandler {
	return &BookHandler{
		bookRepo: bookrepo,
	}
}

func (b BookHandler) NewBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err := b.bookRepo.Save(book)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (b BookHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	books, err := b.bookRepo.Get(int64(id))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, books)
}
