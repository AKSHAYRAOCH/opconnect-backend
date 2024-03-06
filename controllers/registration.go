package controllers

import (
	"context"
	"fmt"
	"net/http"
	"opconnect-backend/db/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm string `json:"confirmpassword"`
}

func Registration(c echo.Context) error {
	Register := new(User)
	
	if err := c.Bind(&Register); err != nil {
		return err
	}
	query:= `insert into Users (Id,Email,Password,Role) values (@Username,@Email,@Password,@Role)`
	values:= pgx.NamedArgs{
		"Username":Register.Username,
		"Email":Register.Email,
		"Password":Register.Password,
		"Role":"Student"
	}
	postgres.DB.Exec(context.Background(),)
	
	return c.String(http.StatusOK,"registered sucessfully")
}
