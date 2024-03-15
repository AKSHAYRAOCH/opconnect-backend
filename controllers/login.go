package controllers

import (
	"context"
	"net/http"
	"opconnect-backend/db/postgres"
	"opconnect-backend/auth"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type userlogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userdetails struct {
	Username string
	Email    string
	Password string
	Role     string
}

func Login(c echo.Context) error {
	loginrequest := new(userlogin)

	if err := c.Bind(loginrequest); err != nil {
		return echo.NewHTTPError(http.StatusUnsupportedMediaType, "invalid json")
	}
	if err := c.Validate(loginrequest); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid body")
	}

	querystring := `select id,email,password,role from users where email=@Email`
	values := pgx.NamedArgs{
		"Email": loginrequest.Email,
	}
	user := new(userdetails)
	if err := postgres.DB.QueryRow(context.Background(), querystring, values).Scan(&user.Username, &user.Email, &user.Password, &user.Role); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if auth.ComparePassword(loginrequest.Password,user.Password){
		return echo.NewHTTPError(http.StatusBadRequest,"incorrect email/password")
	}
	jwt,err:= auth.Generatejwt(user.Username,user.Role)
	if err!=nil{
		return echo.NewHTTPError(http.StatusInternalServerError,"unable to generate jwt")
	}
	return c.JSON(http.StatusOK,jwt)
}
