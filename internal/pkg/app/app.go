package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-http-server/internal/app/endpoint"
	"go-http-server/internal/app/middleware"
	"go-http-server/internal/app/service"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	server   *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.service = service.New()
	a.endpoint = endpoint.New(a.service)
	a.server = echo.New()
	a.server.Use(middleware.RoleCheck)
	a.server.GET("/status", a.endpoint.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")
	err := a.server.Start(":8081")
	if err != nil {
		a.server.Logger.Fatal(err)
	}
	return nil
}
