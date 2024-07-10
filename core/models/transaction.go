package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Title           string
	Amount          uint
	Status          bool
	Type            string
	TransactionDate time.Time
	Description     string
	Reference       string
	WalletID        uint
}
