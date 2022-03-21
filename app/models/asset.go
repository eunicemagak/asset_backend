package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	ID           uint      `gorm:"id"`
	Title        string    `json:"title"`
	SerialNumber string    `json:"serialnumber"`
	Description  string    `json:"description"`
	Price        string    `json:"price"`
	ImageID      uint      `json:"image_id"`
	ImageType    string    `json:"image_type"`
	IsAssigned   bool      `json:"isAssigned" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (asset *Asset) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Asset{}).Count(&total)

	return total
}

func (asset *Asset) Take(db *gorm.DB, limit int, offset int) interface{} {
	var assets []Asset

	db.Offset(offset).Limit(limit).Preload("Images").Find(&assets)

	return assets
}
