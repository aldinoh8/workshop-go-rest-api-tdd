package entity

type Task struct {
	ID          uint
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
