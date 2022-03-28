package models

import (
	"time"
)

type Accesorie struct {
	ID           uint      `gorm:"id"`
	Title        string    `json:"title"`
	SerialNumber string    `json:"serialnumber"`
	Description  string    `json:"description"`
	IsAssigned   bool      `json:"is_assigned" gorm:"default:false"`
	ImageID      uint      `json:"image_id"`
	ImageType    string    `json:"image_type"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
