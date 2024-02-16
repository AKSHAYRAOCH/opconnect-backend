package main

import (
	"net/http"
	"opconnect-backend/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/register",controllers.Registration)
	e.POST("/login",controllers.Login)
	e.Logger.Fatal(e.Start(":1323"))
}
