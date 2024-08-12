package database

import (
	"log"

	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conn *gorm.DB

type DB struct {
	gorm *gorm.DB
}

func NewDB() *DB {
	return &DB{gorm: conn}
}

func (db *DB) CreateNote(note *models.Note) error {
	return db.gorm.Create(note).Error
}

func (db *DB) GetNoteByKey(key string) (*models.Note, error) {
	var note models.Note
	err := db.gorm.Where("unique_key = ?", key).First(&note).Error
	return &note, err
}

func (db *DB) UpdateNoteViews(note *models.Note) error {
	return db.gorm.Save(note).Error
}

func (db *DB) GetUserNotes(username string) ([]models.Note, error) {
	var notes []models.Note
	err := db.gorm.Where("username = ?", username).Find(&notes).Error
	return notes, err
}

func (db *DB) CreateUser(user *models.User) error {
	return db.gorm.Create(&user).Error
}

func (db *DB) GetUserByUsernameAndPassword(username, password string) error {
	var user models.User
	return db.gorm.Where("username = ? AND password = ?", username, password).First(&user).Error
}

func (db *DB) ConnectDatabase() {

	var err error
	db.gorm, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.gorm.AutoMigrate(&models.Note{}, &models.User{})
}
