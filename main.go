package main

import (
	"encoding/json"
	"fmt"
	"stranger-words/config"
	"stranger-words/models"
)

func main() {
	db := config.GetDB()
	var words []models.Word
	// author := models.Author{
	// 	Name:  "Hundera",
	// 	Email: "test@test.com",
	// }
	// word := models.Word{
	// 	Title:  "Title",
	// 	Body:   "body",
	// 	Author: author,
	// }
	// db.Create(&word)
	db.Preload("Author").Find(&words)
	for _, word := range words {
		wordJson, err := json.Marshal(word)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(wordJson))
	}
}
