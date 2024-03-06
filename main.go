package main

import (
	"opconnect-backend/controllers"
	"opconnect-backend/db/postgres"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func main() {
	godotenv.Load()
	e := echo.New()
	postgres.DBConnect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/register", controllers.Registration)
	e.POST("/login", controllers.Login)
	g := e.Group("/api")
	g.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	
	e.Logger.Fatal(e.Start(":1323"))
}
