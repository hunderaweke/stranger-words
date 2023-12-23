package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stranger-words/config"
	"stranger-words/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var db = config.GetDB()

func GetWords(w http.ResponseWriter, r *http.Request) {
	words := models.GetWords(db)
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
	word := models.GetWord(db, id)
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
	word := models.CreateWord(db, newWord)
	json, err := json.Marshal(word)
	if err != nil {
		panic(err)
	}
	w.Write(json)
}
