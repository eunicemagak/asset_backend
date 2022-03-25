package models

// import (
// 	"gorm.io/gorm"
// )

// type Assign struct {
// 	gorm.Model
// 	ID uint `gorm:"primarykey"`

// 	Asset   Asset `json:"asset" gorm:"foreignKey:AssignID;references:ID"`
// 	AssetID uint  `json:"assetID"`
// 	User    User  `json:"user" gorm:"foreignKey:AssignID;references:ID"`
// 	UserID  uint  `json:"assignID"`
// }

// func (assign *Assign) Count(db *gorm.DB) int64 {
// 	var total int64
// 	db.Model(&Assign{}).Count(&total)

// 	return total
// }

// func (assign *Assign) Take(db *gorm.DB, limit int, offset int) interface{} {
// 	var assets []Asset
// 	var users []User

// 	data := db.Offset(offset).Limit(limit).Find(&users, &assets)

// 	return data
// }
