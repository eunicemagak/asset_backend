package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`

	DepartmentID uint `json:"department_id"`
	// AccesorieID  uint `json:"accessorie_id"`
	// Departments []Department `json:"departments" gorm:"many2many:user_departments;"`
	Assets []Asset `json:"assets" gorm:"many2many:user_assets;"`

	Department Department `json:"department" `
	// Accessory  Accessory  `json:"accessory"`

	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)

	return total
}

//Relationship btwn a User and Asset
func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User

	db.Preload("Assets").Preload("Department").Offset(offset).Limit(limit).Find(&users)

	return users
}
