package models

import (
	"gorm.io/gorm"
)

type Status struct {
	ID     uint   `gorm:"id"`
	Status string `json:"status"`
}

func (status *Status) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Tag{}).Count(&total)

	return total
}

func (status *Status) Take(db *gorm.DB, limit int, offset int) interface{} {
	var tags []Tag

	db.Offset(offset).Limit(limit).Find(&tags)

	return tags
}
