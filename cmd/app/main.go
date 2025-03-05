package main

import (
	"log"
	"ruchka/internal/database"
	"ruchka/internal/handlers"
	"ruchka/internal/taskService"
	"ruchka/internal/userService"
	"ruchka/internal/web/tasks"
	"ruchka/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Репозитории и сервисы для задач
	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	// Репозитории и сервисы для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandlers(userService)

	// Инициализация Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация хендлеров для задач
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	// Регистрация хендлеров для пользователей
	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
