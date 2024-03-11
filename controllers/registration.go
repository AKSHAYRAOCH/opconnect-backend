package controllers

import (
	"context"
	"net/http"
	"opconnect-backend/db/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Confirm  string `json:"confirmpassword" validate:"required"`
}

func Registration(c echo.Context) error {
	register := new(User)

	if err := c.Bind(register); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid json")
	}
	if err := c.Validate(register); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid body")
	}
	query := `insert into Users (Id,Email,Password,Role) values (@Username,@Email,@Password,@Role)`
	values := pgx.NamedArgs{
		"Username": register.Username,
		"Email":    register.Email,
		"Password": register.Password,
		"Role":     "Student",
	}

	_, err := postgres.DB.Exec(context.Background(), query, values)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "account already exist's")
	}

	return c.String(http.StatusOK, "registered sucessfully")
}
