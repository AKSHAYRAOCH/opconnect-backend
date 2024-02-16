package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Registration(c echo.Context) error{
	return c.String(http.StatusOK,"this is the registration controller")
}