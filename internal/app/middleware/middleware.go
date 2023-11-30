package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRole := ctx.Request().Header.Get("User-Role")

		if userRole == "admin" {
			log.Println("red button user detected")
		}

		return next(ctx)
	}
}
