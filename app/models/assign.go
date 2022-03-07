package models

import (
	"time"

	"gorm.io/gorm"
)

type Assign struct {
	ID        uint  `gorm:"primarykey"`
	UserId    uint  `json:"user_id"`
	AssetId   uint  `json:"asset_id"`
	Asset     Asset `json:"asset"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `json:"user"`
}

func (assign *Assign) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Assign{}).Count(&total)

	return total
}

func (assign *Assign) Take(db *gorm.DB, limit int, offset int) interface{} {
	var assets []Asset

	db.Offset(offset).Limit(limit).Find(&assets)

	return assets
	var users []User

	db.Offset(offset).Limit(limit).Find(&users)

	return users
}
