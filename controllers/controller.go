package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stranger-words/config"
	"stranger-words/models"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

var db = config.GetDB()

func GetWords(w http.ResponseWriter, r *http.Request) {
	var words []models.Word
	if err := db.Preload("Author").Find(&words).Error; err != nil {
		panic(err)
	}
	json, err := json.Marshal(words)
	if err != nil {
		panic(err)
	}
	w.Write(json)
}

func GetWord(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	var word models.Word
	if err := db.Preload("Author").First(&word, id).Error; err != nil {
		panic(err)
	}
	json, err := json.Marshal(word)
	if err != nil {
		panic(err)
	}
	w.Write(json)
}
func CreateWord(w http.ResponseWriter, r *http.Request) {
	var newWord models.Word
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), &newWord); err != nil {
			return
		}
	}
	newAuthor := newWord.Author
	existingAuthor, err := createOrUpdateAuthor(&newAuthor)
	if err != nil {
		log.Fatal(err)
	}
	var author models.Author
	if existingAuthor != nil {
		author = models.Author(*existingAuthor)
	} else {
		author = newAuthor
	}
	word := models.Word{
		Title:  newWord.Title,
		Body:   newWord.Body,
		Author: author,
	}
	db.Create(&word)
	json, err := json.Marshal(newWord)
	if err != nil {
		panic(err)
	}
	w.Write(json)
}
func createOrUpdateAuthor(author *models.Author) (*models.Author, error) {
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
