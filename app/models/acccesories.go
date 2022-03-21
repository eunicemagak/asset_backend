package models

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// type Accessory struct {
// 	ID           uint           `gorm:"id"`
// 	Title        string         `json:"title"`
// 	SerialNumber string         `json:"serialnumber"`
// 	Description  string         `json:"description"`
// 	Image        string         `json:"image"`
// 	Price        string         `json:"price"`
// 	CreatedAt    time.Time      `gorm:"index"`
// 	UpdatedAt    time.Time      `gorm:"index"`
// 	DeletedAt    gorm.DeletedAt `gorm:"index"`
// }

// func (acccesorie *Accessory) Count(db *gorm.DB) int64 {
// 	var total int64
// 	db.Model(&Accessory{}).Count(&total)

// 	return total
// }

// func (accessory *Accessory) Take(db *gorm.DB, limit int, offset int) interface{} {
// 	var accessories []Accessory

// 	db.Offset(offset).Limit(limit).Find(&accessories)

// 	return accessories
// }
