package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primarykey"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  []byte         `json:"-"`
	RoleId    uint           `json:"role_id"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (admin *Admin) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	admin.Password = hashedPassword
}

func (admin *Admin) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
}

func (admin *Admin) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Admin{}).Count(&total)

	return total
}

func (admin *Admin) Take(db *gorm.DB, limit int, offset int) interface{} {
	var admins []Admin

	db.Preload("Role").Offset(offset).Limit(limit).Find(&admins)

	return admins
}
