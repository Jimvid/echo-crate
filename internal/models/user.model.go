package models

type User struct {
	ID     int    `gorm:"primarkey;size:16"`
	UserID string `gorm:"unique;id"`
	Email  string `gorm:"unique;email"`
}
