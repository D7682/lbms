package main

import (
	"encoding/json"
	"fmt"
	"lbms/app/models"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	books := []models.Book{
		{
			Title:  "The Alchemist",
			Author: "Paulo Coelho",
		},
		{
			Title:  "The Fault in our Stars (Hardcover)",
			Author: "John Green",
		},
		{
			Title:  "The Hunger Games (The Hunger Games #1)",
			Author: "Suzanne Collins",
		},
		{
			Title:  "To Kill a Mockingbird (Paperback)",
			Author: "Harper Lee",
		},
		{
			Title:  "The Da Vinci Code (Robert Langdon #2)",
			Author: "Dan Brown",
		},
		{
			Title:  "The Perks of Being a Wallflower (Paperback)",
			Author: "Stephen Chbosky",
		},
		{
			Title:  "Harry Potter and the Sorcerer's Stone (Harry Potter, #1)",
			Author: "J.K. Rowling",
		},
		{
			Title:  "Looking for Alaska (Paperback)",
			Author: "John Green",
		},
	}

	start := time.Now()
	for _, val := range books {
		data, err := json.Marshal(val)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.Post("http://localhost:8080/books", "application/json", strings.NewReader(string(data)))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
	}
	fmt.Printf("Elapsed: %v\n", time.Since(start))
}
