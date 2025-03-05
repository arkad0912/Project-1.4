package userService

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id,omitempty" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Password  string         `json:"password,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
