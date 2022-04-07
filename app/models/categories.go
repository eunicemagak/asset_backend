package models

import (
	"time"

	"gorm.io/gorm"
)

type Categorie struct {
	ID        uint      `gorm:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (categorie *Categorie) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Categorie{}).Count(&total)

	return total
}

func (categorie *Categorie) Take(db *gorm.DB, limit int, offset int) interface{} {
	var categories []Categorie

	db.Offset(offset).Limit(limit).Find(&categories)

	return categories
}
