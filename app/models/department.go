package models

import "gorm.io/gorm"

type Department struct {
	ID       uint   `gorm:"id"`
	Title    string `json:"title"`
	UserName string `json:"name"`
}

func (department *Department) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Department{}).Count(&total)

	return total
}

func (department *Department) Take(db *gorm.DB, limit int, offset int) interface{} {
	var departments []Department

	db.Offset(offset).Limit(limit).Find(&departments)

	return departments
}
