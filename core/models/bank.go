package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Name string
	Code string
	Ussd string
	Logo string
	Slug string
}

func (bank *Bank) BeforeCreate(tx *gorm.DB) (err error) {
	bank.Slug = slug.Make(bank.Name)
	return
}
