package userService

import (
	//"time"

	"ruchka/internal/taskService"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint               `json:"id,omitempty" gorm:"primaryKey"`
	Email    string             `json:"email"`
	Password string             `json:"password,omitempty"`
	Tasks    []taskService.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"` // Используйте taskService.Task
	// CreatedAt time.Time      `json:"created_at,omitempty"`
	// UpdatedAt time.Time      `json:"updated_at,omitempty"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"

}
