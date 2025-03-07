package userService

import "ruchka/internal/taskService"

// UserService предоставляет методы для работы с пользователями.
type UserService struct {
	repo UserRepository
}

// NewUserService создает новый экземпляр UserService.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser создает нового пользователя.
func (s *UserService) CreateUser(user *User) (*User, error) {
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers возвращает список всех пользователей.
func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

// GetUserByID возвращает пользователя по его ID.
func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

// UpdateUser обновляет данные пользователя по его ID.
func (s *UserService) UpdateUser(id uint, user *User) (*User, error) {
	err := s.repo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	// Получаем обновленного пользователя
	updatedUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// DeleteUser удаляет пользователя по его ID.
func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	// Реализация метода для получения задач пользователя
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user.Tasks, nil
}
