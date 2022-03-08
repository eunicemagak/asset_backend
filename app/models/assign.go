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

	ID        uint           `gorm:"primarykey"`
	UserId    uint           `json:"user_id"`
	AssetId   uint           `json:"asset_id"`
	User      User           `json:"user"`
	Asset     Asset          `json:"asset"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// func (assign *Assign) Take(db *gorm.DB, limit int, offset int) interface{} {
// 	var assets []Asset
// 	var users []User

// 	data := db.Offset(offset).Limit(limit).Find(&users, &assets)

// 	return data
// }
