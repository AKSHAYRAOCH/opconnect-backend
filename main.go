package main

import (
	"opconnect-backend/controllers"
	"opconnect-backend/db/postgres"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	godotenv.Load()
	postgres.DBConnect()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/register", controllers.Registration)
	e.POST("/login", controllers.Login)
	g := e.Group("/api")
	g.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
