package models

type User struct {
	Username string `gorm:"uniqueIndex"`
	Password string
	Notes    []Note `gorm:"foreignKey:Username;references:Username"`
}
