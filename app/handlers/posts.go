package handlers

import (
	"encoding/json"
	"github.com/vavilen84/go_skills_up_project/models"
	"github.com/vavilen84/go_skills_up_project/store"
	"log"
	"net/http"
)

type PostsController struct{}

func (c *PostsController) Create(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	dec := json.NewDecoder(r.Body)
	model := models.Post{}
	err := dec.Decode(&model)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = model.Create(db)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
