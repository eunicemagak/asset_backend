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
	IsAssigned   bool      `json:"is_assigned" gorm:"default:false"`
	IsClearedOf  bool      `json:"is_cleared_of" gorm:"default:false"`
	IsDamaged    bool      `json:"is_damaged" gorm:"default:false"`
	PurchaseDate string    `json:"purchase_date"`
	AssignedTo   string    `json:"assigned_to"`
	ImageID      uint      `json:"image_id"`
	ImageType    string    `json:"image_type"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Categories []Categorie `json:"categorie" gorm:"many2many:accesorie_categorie;"`
	Statuses   []Status    `json:"status" gorm:"many2many:accesorie_status;"`
}

func (accesorie *Accesorie) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Accesorie{}).Count(&total)

	return total
}

func (acccesorie *Accesorie) Take(db *gorm.DB, limit int, offset int) interface{} {
	var acccesories []Accesorie

	db.Offset(offset).Limit(limit).Preload("Categories").Preload("Statuses").Find(&acccesories)

	return acccesories
}
