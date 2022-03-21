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
	AccesorieID  uint `json:"accessorie_id"`

	Assets     []Asset    `json:"assets" gorm:"many2many:user_assets;"`
	Department Department `json:"department"`
	Accesorie  Accesorie  `json:"accesorie"`

	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
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

<<<<<<< HEAD
	// db.Offset(offset).Limit(limit).Find(&users)
=======
>>>>>>> 2fe1552807d5a5c090e33e8e4898e3f5753702b8
	db.Preload("Assets").Preload("Department").Preload("Accesorie").Offset(offset).Limit(limit).Find(&users)

	return users
}
