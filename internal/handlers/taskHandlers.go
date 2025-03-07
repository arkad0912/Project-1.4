package handlers

import (
	"context"

	//"strconv"

	"ruchka/internal/taskService"
	"ruchka/internal/web/tasks"
)

type TaskHandler struct {
	taskService *taskService.TaskService
}

func NewTaskHandler(taskService *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.taskService.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
			//DeletedAt: &tsk.DeletedAt.Time, // Добавляем поле DeletedAt
		}

		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.taskService.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *TaskHandler) GetUsersUserIdTasks(_ context.Context, request tasks.GetUsersUserIdTasksRequestObject) (tasks.GetUsersUserIdTasksResponseObject, error) {
	userID := request.UserId

	// Получаем задачи для пользователя
	userTasks, err := h.taskService.GetTasksByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	response := tasks.GetUsersUserIdTasks200JSONResponse{}
	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
			//DeletedAt: &tsk.DeletedAt.Time,
		}
		if tsk.DeletedAt.Valid {
			task.DeletedAt = &tsk.DeletedAt.Time
		}

		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksTaskId(ctx context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
	// Получаем ID задачи из запроса
	taskID := request.TaskId

	// Получаем обновленные данные из тела запроса
	taskUpdates := request.Body

	// Преобразуем обновленные данные в структуру taskService.Task
	updatedTask := taskService.Task{
		Task:   *taskUpdates.Task,
		IsDone: *taskUpdates.IsDone,
	}

	// Вызываем метод сервиса для обновления задачи
	updatedTaskResult, err := h.taskService.UpdateTaskByID(uint(taskID), updatedTask)
	if err != nil {
		return nil, err
	}

	// Создаем структуру ответа
	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &updatedTaskResult.ID,
		Task:   &updatedTaskResult.Task,
		IsDone: &updatedTaskResult.IsDone,
	}

	// Возвращаем ответ
	return response, nil
}
func (h *TaskHandler) DeleteTasksTaskId(ctx context.Context, request tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	// Получаем ID задачи из запроса
	taskID := request.TaskId

	// Вызываем метод сервиса для удаления задачи
	deletedAt, err := h.taskService.DeleteTaskByID(uint(taskID))
	if err != nil {
		return nil, err
	}

	// Возвращаем дату удаления в ответе
	response := tasks.DeleteTasksTaskId200JSONResponse{
		DeletedAt: &deletedAt,
	}

	return response, nil
}
