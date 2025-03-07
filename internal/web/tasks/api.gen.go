// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	// CreatedAt Timestamp when the task was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// DeletedAt Timestamp when the task was deleted (nullable)
	DeletedAt *time.Time `json:"deleted_at"`
	Id        *uint      `json:"id,omitempty"`
	IsDone    *bool      `json:"is_done,omitempty"`
	Task      *string    `json:"task,omitempty"`

	// UpdatedAt Timestamp when the task was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UserId    *uint      `json:"user_id,omitempty"`
}

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = Task

// PatchTasksTaskIdJSONRequestBody defines body for PatchTasksTaskId for application/json ContentType.
type PatchTasksTaskIdJSONRequestBody = Task

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx echo.Context) error
	// Delete a task by ID
	// (DELETE /tasks/{taskId})
	DeleteTasksTaskId(ctx echo.Context, taskId uint) error
	// Update a task by ID
	// (PATCH /tasks/{taskId})
	PatchTasksTaskId(ctx echo.Context, taskId uint) error
	// Get tasks by user ID
	// (GET /users/{userId}/tasks)
	GetUsersUserIdTasks(ctx echo.Context, userId uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PostTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasks(ctx)
	return err
}

// DeleteTasksTaskId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTasksTaskId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "taskId" -------------
	var taskId uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskId", runtime.ParamLocationPath, ctx.Param("taskId"), &taskId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter taskId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTasksTaskId(ctx, taskId)
	return err
}

// PatchTasksTaskId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchTasksTaskId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "taskId" -------------
	var taskId uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskId", runtime.ParamLocationPath, ctx.Param("taskId"), &taskId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter taskId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchTasksTaskId(ctx, taskId)
	return err
}

// GetUsersUserIdTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersUserIdTasks(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersUserIdTasks(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.POST(baseURL+"/tasks", wrapper.PostTasks)
	router.DELETE(baseURL+"/tasks/:taskId", wrapper.DeleteTasksTaskId)
	router.PATCH(baseURL+"/tasks/:taskId", wrapper.PatchTasksTaskId)
	router.GET(baseURL+"/users/:userId/tasks", wrapper.GetUsersUserIdTasks)

}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksRequestObject struct {
	Body *PostTasksJSONRequestBody
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

type PostTasks201JSONResponse Task

func (response PostTasks201JSONResponse) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTasksTaskIdRequestObject struct {
	TaskId uint `json:"taskId"`
}

type DeleteTasksTaskIdResponseObject interface {
	VisitDeleteTasksTaskIdResponse(w http.ResponseWriter) error
}

type DeleteTasksTaskId200JSONResponse struct {
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (response DeleteTasksTaskId200JSONResponse) VisitDeleteTasksTaskIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTasksTaskId404Response struct {
}

func (response DeleteTasksTaskId404Response) VisitDeleteTasksTaskIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchTasksTaskIdRequestObject struct {
	TaskId uint `json:"taskId"`
	Body   *PatchTasksTaskIdJSONRequestBody
}

type PatchTasksTaskIdResponseObject interface {
	VisitPatchTasksTaskIdResponse(w http.ResponseWriter) error
}

type PatchTasksTaskId200JSONResponse Task

func (response PatchTasksTaskId200JSONResponse) VisitPatchTasksTaskIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchTasksTaskId404Response struct {
}

func (response PatchTasksTaskId404Response) VisitPatchTasksTaskIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetUsersUserIdTasksRequestObject struct {
	UserId uint `json:"userId"`
}

type GetUsersUserIdTasksResponseObject interface {
	VisitGetUsersUserIdTasksResponse(w http.ResponseWriter) error
}

type GetUsersUserIdTasks200JSONResponse []Task

func (response GetUsersUserIdTasks200JSONResponse) VisitGetUsersUserIdTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error)
	// Delete a task by ID
	// (DELETE /tasks/{taskId})
	DeleteTasksTaskId(ctx context.Context, request DeleteTasksTaskIdRequestObject) (DeleteTasksTaskIdResponseObject, error)
	// Update a task by ID
	// (PATCH /tasks/{taskId})
	PatchTasksTaskId(ctx context.Context, request PatchTasksTaskIdRequestObject) (PatchTasksTaskIdResponseObject, error)
	// Get tasks by user ID
	// (GET /users/{userId}/tasks)
	GetUsersUserIdTasks(ctx context.Context, request GetUsersUserIdTasksRequestObject) (GetUsersUserIdTasksResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(ctx echo.Context) error {
	var request GetTasksRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx.Request().Context(), request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		return validResponse.VisitGetTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTasks operation middleware
func (sh *strictHandler) PostTasks(ctx echo.Context) error {
	var request PostTasksRequestObject

	var body PostTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasks(ctx.Request().Context(), request.(PostTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTasksResponseObject); ok {
		return validResponse.VisitPostTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTasksTaskId operation middleware
func (sh *strictHandler) DeleteTasksTaskId(ctx echo.Context, taskId uint) error {
	var request DeleteTasksTaskIdRequestObject

	request.TaskId = taskId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTasksTaskId(ctx.Request().Context(), request.(DeleteTasksTaskIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTasksTaskId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTasksTaskIdResponseObject); ok {
		return validResponse.VisitDeleteTasksTaskIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchTasksTaskId operation middleware
func (sh *strictHandler) PatchTasksTaskId(ctx echo.Context, taskId uint) error {
	var request PatchTasksTaskIdRequestObject

	request.TaskId = taskId

	var body PatchTasksTaskIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTasksTaskId(ctx.Request().Context(), request.(PatchTasksTaskIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTasksTaskId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTasksTaskIdResponseObject); ok {
		return validResponse.VisitPatchTasksTaskIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsersUserIdTasks operation middleware
func (sh *strictHandler) GetUsersUserIdTasks(ctx echo.Context, userId uint) error {
	var request GetUsersUserIdTasksRequestObject

	request.UserId = userId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersUserIdTasks(ctx.Request().Context(), request.(GetUsersUserIdTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersUserIdTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersUserIdTasksResponseObject); ok {
		return validResponse.VisitGetUsersUserIdTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
