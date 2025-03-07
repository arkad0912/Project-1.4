package userService

import (
	//"ruchka/internal/taskService"

	"gorm.io/gorm"
)

// UserRepository определяет интерфейс для работы с пользователями в базе данных.
type UserRepository interface {
	CreateUser(user *User) error
	GetUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	UpdateUser(id uint, user *User) error
	DeleteUser(id uint) error
}

// userRepository реализует интерфейс UserRepository.
type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository создает новый экземпляр UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

// CreateUser создает нового пользователя в базе данных.
func (r *userRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

// GetUsers возвращает список всех пользователей из базы данных.
func (r *userRepository) GetUsers() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

// GetUserByID возвращает пользователя по его ID.
func (r *userRepository) GetUserByID(id uint) (*User, error) {
	var user User
	err := r.DB.First(&user, id).Error
	return &user, err
}

// UpdateUser обновляет данные пользователя по его ID.
func (r *userRepository) UpdateUser(id uint, user *User) error {
	return r.DB.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

// DeleteUser удаляет пользователя по его ID.
func (r *userRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&User{}, id).Error
}
