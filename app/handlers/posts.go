package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/vavilen84/go_skills_up_project/helpers"
	"github.com/vavilen84/go_skills_up_project/models"
	"github.com/vavilen84/go_skills_up_project/store"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PostsController struct{}

func (c *PostsController) Create(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	dec := json.NewDecoder(r.Body)
	model := models.Post{}
	err := dec.Decode(&model)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = models.CreatePost(db, &model)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *PostsController) Update(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	idString := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(idString)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	post, err := models.GetOnePostByID(db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	dec := json.NewDecoder(r.Body)
	model := models.Post{}
	err = dec.Decode(&model)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	post.Description = model.Description

	err = models.UpdatePost(db, post)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *PostsController) Delete(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	idString := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(idString)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	post, err := models.GetOnePostByID(db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	err = models.DeletePost(db, post)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *PostsController) GetOne(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	idString := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(idString)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	post, err := models.GetOnePostByID(db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(post)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *PostsController) GetAll(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	posts, err := models.GetAllPosts(db)
	b, err := json.Marshal(posts)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
