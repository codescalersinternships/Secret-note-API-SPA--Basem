package models

import "time"

type Note struct {
	ID         uint   `gorm: "primary_key"`
	Content    string `gorm: "type:text"`
	UniqueKey  string `gorm: "uniqueIndex"`
	Expiration time.Time
	CreatedAt  time.Time `gorm: "autoCreateTime"`
	MaxViews   int
	Views      int
	UserID     uint
}
