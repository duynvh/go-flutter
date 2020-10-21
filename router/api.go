package router

import (
	"github.com/labstack/echo/v4"
	"go-flutter/handler"
	"go-flutter/middleware"
)

type API struct {
	Echo        *echo.Echo
	AuthHandler handler.AuthHandler
	RepoHandler handler.RepoHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/auth/sign-in", api.AuthHandler.HandleSignIn)
	api.Echo.POST("/auth/sign-up", api.AuthHandler.HandleSignUp)

	user := api.Echo.Group("/users", middleware.JWTMiddleware())
	user.GET("/profile", api.AuthHandler.Profile)
	user.PUT("/profile/update", api.AuthHandler.UpdateProfile)

	// github repo
	github := api.Echo.Group("/github", middleware.JWTMiddleware())
	github.GET("/trending", api.RepoHandler.RepoTrending)

	// bookmark
	bookmark := api.Echo.Group("/bookmark", middleware.JWTMiddleware())
	bookmark.GET("/list", api.RepoHandler.SelectBookmarks)
	bookmark.POST("/add", api.RepoHandler.Bookmark)
	bookmark.DELETE("/delete", api.RepoHandler.DelBookmark)
}
