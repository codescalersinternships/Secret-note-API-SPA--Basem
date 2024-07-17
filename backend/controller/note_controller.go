package controllers

import (
	"net/http"
	"time"

	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/database"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/models"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/utils"
	"github.com/gin-gonic/gin"
)

type NoteController struct {
	DB *database.DB
}

func NewNoteController(db *database.DB) *NoteController {
	return &NoteController{DB: db}
}

type CreateNoteInput struct {
	Content    string `json:"content" binding:"required"`
	Expiration string `json:"expiration" binding:"required"`
	MaxViews   int    `json:"maxViews" binding:"required,min=1"`
}

func (nc *NoteController) CreateNote(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		userID = nil
	} else {
		userID = userID.(string)
	}

	var input CreateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expirationTime, err := time.Parse(time.RFC3339, input.Expiration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration date"})
		return
	}

	if input.MaxViews < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Max views must be greater than 0"})
		return
	}

	note := models.Note{
		Content:    input.Content,
		UniqueKey:  utils.GenerateUniqueKey(),
		Expiration: expirationTime,
		MaxViews:   input.MaxViews,
		Views:      0,
		Username:   userID.(string),
	}

	nc.DB.CreateNote(&note)
	c.JSON(http.StatusOK, gin.H{"url": "/note/" + note.UniqueKey})
}

func (nc *NoteController) GetNoteByKey(c *gin.Context) {
	uniqueKey := c.Param("key")

	note, err := nc.DB.GetNoteByKey(uniqueKey)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if note.Expiration.Before(time.Now()) || note.Views >= note.MaxViews {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note has expired or has reached Max views "})
		return
	}

	note.Views++
	nc.DB.UpdateNoteViews(note)
	c.JSON(http.StatusOK, gin.H{"content": note.Content})
}

func (nc *NoteController) GetUserNotes(c *gin.Context) {

	userID, _ := c.Get("userID")

	notes, err := nc.DB.GetUserNotes(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Notes not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notes": notes})

}
