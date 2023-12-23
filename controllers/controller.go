package controllers

import (
	"encoding/json"
	"net/http"
	"stranger-words/config"
	"stranger-words/models"
	"strconv"

	"github.com/go-chi/chi/v5"
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
