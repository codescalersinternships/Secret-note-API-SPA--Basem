package routes

import (
	controllers "github.com/codescalersinternships/Secret-note-API-SPA--Basem/controller"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/note", utils.OptionalAuthMiddleware(), controllers.CreateNote)
	r.GET("/note/:key", controllers.GetNoteByKey)

	auth := r.Group("/")
	auth.Use(utils.AuthMiddleware())
	{

		auth.GET("/user/notes", controllers.GetUserNotes)
	}
}
