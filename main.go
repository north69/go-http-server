package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello")

	server := echo.New()

	server.Use(Middleware)

	server.GET("/status", Handler)

	server.Logger.Fatal(server.Start(":8081"))
}

func Handler(c echo.Context) error {
	d := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	s := fmt.Sprintf("Days remained: %d", int64(dur.Hours()/24))
	return c.String(http.StatusOK, s)
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRole := ctx.Request().Header.Get("User-Role")

		if userRole == "admin" {
			log.Println("red button user detected")
		}

		return next(ctx)
	}
}
