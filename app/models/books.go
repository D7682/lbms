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

/* func (b BookHandler) NewBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err := b.driver.Write("books", book.Title, book)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (b BookHandler) GetAll(c *gin.Context) {
	books, err := b.driver.ReadAll("books")
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	var bookList []Book
	for _, val := range books {
		var book Book
		err = json.Unmarshal([]byte(val), &book)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		bookList = append(bookList, book)
	}
	c.JSON(http.StatusOK, bookList)
} */
