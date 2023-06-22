package handler

import (
	"net/http"
	"strconv"
	"workshoptdd/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h Handler) CreateTask(ctx *gin.Context) {
	var newTask entity.Task

	err := ctx.ShouldBindJSON(&newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed create task",
		})
		return
	}

	h.DB.Create(&newTask)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success create task",
	})
}

func (h Handler) GetTasks(ctx *gin.Context) {
	var tasks []entity.Task

	h.DB.Find(&tasks)

	ctx.JSON(http.StatusOK, tasks)
}

func (h Handler) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id invalid",
		})
		return
	}

	h.DB.Delete(&entity.Task{}, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success delete task",
	})
}
