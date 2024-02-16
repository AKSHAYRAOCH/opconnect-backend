package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


func Login(c echo.Context) error {
	return c.String(http.StatusOK,"this is the Login handler")
}