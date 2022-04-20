package models

type Status struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}

// func (state *Status) Count(db *gorm.DB) int64 {
// 	var total int64
// 	db.Model(&Status{}).Count(&total)

// 	return total
// }

// func (state *Status) Take(db *gorm.DB, limit int, offset int) interface{} {
// 	var status []Status

// 	db.Offset(offset).Limit(limit).Find(&status)

// 	return status
// }
