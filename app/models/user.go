package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Assets    []Asset   `json:"assets" gorm:"many2many:user_assets;"`

	// Tags        []Tag        `json:"tag" gorm:"many2many:user_tags;"`
	Accesories  []Accesorie  `json:"accesories" gorm:"many2many:user_accesories;"`
	Departments []Department `json:"departments" gorm:"many2many:user_department;"`
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)

	return total
}

//Relationship btwn a User and Asset
func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User

	db.Preload("Assets").Preload("Accesories").Preload("Departments").Offset(offset).Limit(limit).Find(&users)

	return users
}
