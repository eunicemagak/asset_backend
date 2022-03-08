package models

import (
	"time"

	"gorm.io/gorm"
)

type Accesorie struct {
	ID           uint           `gorm:"id"`
	Title        string         `json:"title"`
	SerialNumber string         `json:"serialnumber"`
	Description  string         `json:"description"`
	Image        string         `json:"image"`
	Price        string         `json:"price"`
	CreatedAt    time.Time      `gorm:"index"`
	UpdatedAt    time.Time      `gorm:"index"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (acccesorie *Accesorie) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Accesorie{}).Count(&total)

	return total
}

func (acccesorie *Accesorie) Take(db *gorm.DB, limit int, offset int) interface{} {
	var acccesories []Accesorie

	db.Offset(offset).Limit(limit).Find(&acccesories)

	return acccesories
}
