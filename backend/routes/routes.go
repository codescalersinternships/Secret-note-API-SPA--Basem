package routes

import (
	controllers "github.com/codescalersinternships/Secret-note-API-SPA--Basem/controller"
	_ "github.com/codescalersinternships/Secret-note-API-SPA--Basem/docs"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine, noteController *controllers.NoteController, userController *controllers.UserController) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
