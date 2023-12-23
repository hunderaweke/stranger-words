package main

import (
	"net/http"
	"stranger-words/models"
	"stranger-words/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func createOrUpdateAuthor(db *gorm.DB, author *models.Author) (*models.Author, error) {
	var existingAuthor models.Author
	result := db.Where("email = ?", author.Email).First(&existingAuthor)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(author).Error; err != nil {
				return nil, err
			}
			return author, nil
		}
		return nil, result.Error
	}
	return &existingAuthor, nil
}

func main() {
	// db := config.GetDB()
	// newAuthor := &models.Author{
	// 	Name:  "John Doe",
	// 	Email: "john@example.com",
	// }
	// existingAuthor, err := createOrUpdateAuthor(db, newAuthor)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var author models.Author
	// if existingAuthor == nil {
	// 	author = models.Author(*existingAuthor)
	// } else {
	// 	author = *newAuthor
	// }
	// word := models.Word{
	// 	Title:  "THis is a title",
	// 	Body:   "this is a body",
	// 	Author: author,
	// }
	// db.Create(&word)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	routes.RegisterRoutes(router)
	http.ListenAndServe(":8080", router)
}
