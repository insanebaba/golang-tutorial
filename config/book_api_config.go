package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// Book a Structure for book type
type Book struct {
	gorm.Model
	Name       string
	AuthorName string
}

var (
	// Books a placeholder book container
	Books = []Book{
		{Name: "A thousand suns", AuthorName: "Author1"},
		{Name: "A thousand moons", AuthorName: "Author2"},
		{Name: "A thousand planets", AuthorName: "Author1"},
		{Name: "A thousand stones", AuthorName: "Author2"},
		{Name: "A thousand bricks", AuthorName: "Author3"},
	}
)

func getBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.ParseUint(params.ByName("id"), 10, 32)
	if err != nil {
		fmt.Fprintf(writer, "invalid id format")
		return
	}
	var book Book
	DB.First(&book, id)
	if book.ID == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Invalid ID")
		return
	}
	response, _ := json.Marshal(book)
	fmt.Fprintf(writer, string(response))
	return
}

func getBooks(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var books []Book
	DB.Find(&books)
	response, _ := json.Marshal(books)
	fmt.Fprintf(writer, string(response))
}

// ConfigureBookRoutes to configure routes for books
func ConfigureBookRoutes(router *httprouter.Router) {
	router.GET("/books/:id", getBook)
	router.GET("/books", getBooks)
}
