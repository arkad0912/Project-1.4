package handlers

import (
	"context"

	//"strconv"

	"ruchka/internal/taskService"
	"ruchka/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
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
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) PatchTasksTaskId(ctx context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
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
	updatedTaskResult, err := h.Service.UpdateTaskByID(uint(taskID), updatedTask)
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
func (h *Handler) DeleteTasksTaskId(ctx context.Context, request tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	// Получаем ID задачи из запроса
	taskID := request.TaskId

	// Вызываем метод сервиса для удаления задачи
	err := h.Service.DeleteTaskByID(uint(taskID))
	if err != nil {
		return nil, err
	}

	// Возвращаем статус 204 No Content
	return tasks.DeleteTasksTaskId204Response{}, nil
}
