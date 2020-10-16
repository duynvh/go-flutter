package main

import (
	"github.com/labstack/echo/v4"
	"go-flutter/db"
	"go-flutter/handler"
	"net/http"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "123123",
		DbName:   "go-flutter",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/", welcome)

	e.GET("/auth/sign-in", handler.HandleSignIn)
	e.GET("/auth/sign-up", handler.HandleSignUp)
	e.Logger.Fatal(e.Start(":1323"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to my app")
}
