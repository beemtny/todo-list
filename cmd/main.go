package main

import (
	"net/http"
	"todolist/handler"
	"todolist/repositories"
	"todolist/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db := repositories.ConnectPostgres()

	var (
		taskRepo    = repositories.NewTasksRepositoryDB(db)
		taskService = service.NewTaskService(taskRepo)
		taskHandler = handler.NewTaskHandler(taskService)

		userRepo    = repositories.NewUsersRepositoryDB(db)
		userService = service.NewUserService(userRepo)
		userHandler = handler.NewUserHandler(userService)
	)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/tasks", taskHandler.GetTasks)
	e.POST("/task", taskHandler.CreateTask)
	e.PUT("/task/:id", taskHandler.UpdateTask)
	e.DELETE("/task/:id", taskHandler.DeleteTask)

	e.POST("/login", userHandler.Login)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
