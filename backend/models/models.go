package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Word struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	AuthorID  uint      `json:"-"`
	Author    Author    `json:"author,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func GetWord(db *gorm.DB, id int) Word {
	var word Word
	if err := db.Preload("Author").First(&word, id).Error; err != nil {
		panic(err)
	}
	return word
}

func GetWords(db *gorm.DB) []Word {
	var words []Word
	if err := db.Preload("Author").Find(&words).Error; err != nil {
		panic(err)
	}
	return words
}

func CreateWord(db *gorm.DB, w Word) Word {
	newAuthor := w.Author
	existingAuthor, err := createOrUpdateAuthor(db, &newAuthor)
	if err != nil {
		log.Fatal(err)
	}
	var author Author
	if existingAuthor != nil {
		author = Author(*existingAuthor)
	} else {
		author = newAuthor
	}
	word := Word{
		Title:  w.Title,
		Body:   w.Body,
		Author: author,
	}
	db.Create(&word)
	return word
}

func createOrUpdateAuthor(db *gorm.DB, author *Author) (*Author, error) {
	var existingAuthor Author
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
