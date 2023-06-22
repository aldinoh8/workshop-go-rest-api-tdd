package routes

import (
	"net/http"
	"workshoptdd/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	app *gin.Engine
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	app = gin.Default()

	// your app endpoints here
	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	taskHandler := handler.NewHandler(db)
	app.POST("/tasks", taskHandler.CreateTask)
	app.GET("/tasks", taskHandler.GetTasks)
	app.DELETE("/tasks/:id", taskHandler.DeleteTask)

	return app
}
