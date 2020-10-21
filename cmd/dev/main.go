package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-flutter/db"
	"go-flutter/log"
	"go-flutter/repository/repo_impl"
	"go-flutter/router"
	"os"
	"strconv"

	//"github.com/joho/godotenv"
	"go-flutter/handler"
	"go-flutter/helper"
	//"go-flutter/log"
	"time"
)

func init() {
	//Load Env
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	log.InitLogger(false)
}

// @title Github Trending API
// @version 1.0
// @description More
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey jwt
// @in header
// @name Authorization

// @host localhost:1323
// @BasePath /
func main() {
	dbPort, _ :=  strconv.Atoi(os.Getenv("DB_PORT"))
	if dbPort == 0 {
		dbPort = 5432
	}

	dbHost :=  os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbUserName :=  os.Getenv("DB_USERNAME")
	if dbUserName == "" {
		dbUserName = "postgres"
	}

	dbPassword :=  os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "123123"
	}

	dbName :=  os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "go-flutter"
	}

	sql := &db.Sql{
		Host:     dbHost,
		Port:     dbPort,
		UserName: dbUserName,
		Password: dbPassword,
		DbName:   dbName,
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	authHandler := handler.AuthHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}

	api := router.API{
		Echo:        e,
		AuthHandler: authHandler,
		RepoHandler: repoHandler,
	}

	api.SetupRouter()

	go scheduleUpdateTrending(360 * time.Second, repoHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
