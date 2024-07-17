package routes

import (
	controllers "github.com/codescalersinternships/Secret-note-API-SPA--Basem/controller"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, noteController *controllers.NoteController, userController *controllers.UserController) {

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/note", utils.OptionalAuthMiddleware(), noteController.CreateNote)
	r.GET("/note/:key", noteController.GetNoteByKey)

	auth := r.Group("/")
	auth.Use(utils.AuthMiddleware())
	{

		auth.GET("/user/notes", noteController.GetUserNotes)
	}
}
