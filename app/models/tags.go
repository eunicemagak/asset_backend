package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint      `gorm:"id"`
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostTag struct {
	TagName string `json:"tag_name"`
}

func (tag *Tag) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Tag{}).Count(&total)
	return total

}

func (tag *Tag) Take(db *gorm.DB, limit int, offset int) interface{} {

	var tags []Tag
	db.Offset(offset).Limit(limit).Find(&tags)

	return tags
}
