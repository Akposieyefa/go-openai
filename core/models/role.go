package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string
	Description string
	Slug        string
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.Slug = slug.Make(role.Name)
	return
}
