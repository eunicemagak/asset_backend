package models

import (
	"time"

	"gorm.io/gorm"
)

type Assign struct {
	gorm.Model

	ID        uint           `gorm:"primarykey"`
	UserId    uint           `json:"user_id"`
	AssetId   uint           `json:"asset_id"`
	User      User           `json:"user"`
	Asset     Asset          `json:"asset"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
}
