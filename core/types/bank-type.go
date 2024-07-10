package types

import (
	"time"

	"gorm.io/gorm"
)

type BankResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Code      string         `json:"code"`
	Ussd      string         `json:"ussd"`
	Logo      string         `json:"logo"`
	Slug      string         `json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
