package main

import (
	"github.com/labstack/echo"
	"go-flutter/db"
	"go-flutter/handler"
	"go-flutter/helper"
	log "go-flutter/log"
	"go-flutter/repository/repo_impl"
	"go-flutter/router"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "go-flutter")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "nguyenduy",
		Password: "123123",
		DbName:   "go-flutter",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	authHandler := handler.AuthHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		AuthHandler: authHandler,
	}

	api.SetupRouter()
	e.Logger.Fatal(e.Start(":1323"))
}
