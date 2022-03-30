package models

import (
	"time"
)

type Department struct {
	ID        uint      `gorm:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func (department *Department) Count(db *gorm.DB) int64 {
// 	var total int64
// 	db.Model(&Department{}).Count(&total)

// 	return total
// }

// func (department *Department) Take(db *gorm.DB, limit int, offset int) interface{} {
// 	var departments []Department

// 	db.Offset(offset).Limit(limit).Find(&departments)

// 	return departments
// }
