package models

import (
	"time"

	"gorm.io/gorm"
)

type Accesorie struct {
	ID           uint      `gorm:"id"`
	Title        string    `json:"title"`
	SerialNumber string    `json:"serialnumber"`
	Description  string    `json:"description"`
	ImageID      uint      `json:"image_id"`
	ImageType    string    `json:"image_type"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (acccesorie *Accesorie) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Accesorie{}).Count(&total)

	return total
}

func (acccesorie *Accesorie) Take(db *gorm.DB, limit int, offset int) interface{} {
	var acccesories []Accesorie

	db.Offset(offset).Limit(limit).Preload("Images").Find(&acccesories)

	return acccesories
}
