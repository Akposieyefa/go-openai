package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Balance      uint
	AcctNumber   string
	Currency     string
	UserID       uint
	Transactions []Transaction
}
