package handlers

import (
	"context"

	"ruchka/internal/userService"
	"ruchka/internal/web/users"
)

type UserHandlers struct {
	userService *userService.UserService
}

func NewUserHandlers(userService *userService.UserService) *UserHandlers {
	return &UserHandlers{userService: userService}
}

// GetUsers возвращает список всех пользователей
func (h *UserHandlers) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.userService.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
			// CreatedAt: &usr.CreatedAt,
			// UpdatedAt: &usr.UpdatedAt,
			// DeletedAt: &usr.DeletedAt.Time,
		}

		response = append(response, user)
	}

	return response, nil
}

// PostUsers создает нового пользователя
func (h *UserHandlers) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {

	user := &userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}

	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
		// CreatedAt: &createdUser.CreatedAt,
		// UpdatedAt: &createdUser.UpdatedAt,
		// DeletedAt: &createdUser.DeletedAt.Time,
	}

	return response, nil
}

// PatchUsersUserId обновляет данные пользователя по его ID
func (h *UserHandlers) PatchUsersUserId(_ context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {

	user := &userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}

	updatedUser, err := h.userService.UpdateUser(request.UserId, user)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersUserId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
		// CreatedAt: &updatedUser.CreatedAt,
		// UpdatedAt: &updatedUser.UpdatedAt,
		// DeletedAt: &updatedUser.DeletedAt.Time,
	}

	return response, nil
}

// DeleteUsersUserId удаляет пользователя по его ID
func (h *UserHandlers) DeleteUsersUserId(_ context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	err := h.userService.DeleteUser(request.UserId)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersUserId204Response{}, nil
}
