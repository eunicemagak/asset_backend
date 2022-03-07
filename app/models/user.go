package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`

	DepartmentID uint      `json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Assets     []Asset      `json:"assets" gorm:"many2many:user_assets;"`
	Department []Department `json:"department" gorm:"many2many" `
}

// func (user *User) SetPassword(password string) {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	user.Password = hashedPassword
// }

// func (user *User) ComparePassword(password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// }

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)

	return total
}

//Relationship btwn a User and Asset
func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User

	db.Offset(offset).Limit(limit).Find(&users)

	return users
}
