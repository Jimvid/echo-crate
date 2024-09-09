package models

type Todo struct {
	ID        int    `gorm:"primarkey;size:16"`
	Title     string `gorm:"title"`
	Completed bool   `gorm:"bool"`
}
