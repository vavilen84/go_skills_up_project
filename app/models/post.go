package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
)

type Post struct {
	Id      string `json:"id" column:"id" validate:"required,uuid4" skip_on_update:"true"`
	Title   string `json:"title" column:"title" validate:"required,min=2,max=255"`
	Content string `json:"content" column:"content" validate:"required,min=2,max=5000"`
}

func (Post) TableName() string {
	return "posts"
}

func (m *Post) Create(db *gorm.DB) error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		log.Println(err)
		return err
	}
	db.Create(m)
	return nil
}
