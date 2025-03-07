package taskService

import (
	//"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task      string         `json:"task"`    // Наш сервер будет ожидать json c полем text
	IsDone    bool           `json:"is_done"` // В GO используем CamelCase, в Json - snake
	UserID    uint           `json:"user_id"` // Связь с пользователем
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
