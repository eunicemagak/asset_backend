package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	ID          uint           `gorm:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	Price       float64        `json:"price"`
	CreatedAt   time.Time      `gorm:"index"`
	UpdatedAt   time.Time      `gorm:"index"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (asset *Asset) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Asset{}).Count(&total)

	return total
}

func (asset *Asset) Take(db *gorm.DB, limit int, offset int) interface{} {
	var assets []Asset

	db.Offset(offset).Limit(limit).Find(&assets)

	return assets
}
