package router

import (
	"github.com/labstack/echo"
	"go-flutter/handler"
	"go-flutter/middleware"
)

type API struct {
	Echo        *echo.Echo
	AuthHandler handler.AuthHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/auth/sign-in", api.AuthHandler.HandleSignIn)
	api.Echo.POST("/auth/sign-up", api.AuthHandler.HandleSignUp)

	user := api.Echo.Group("/users", middleware.JWTMiddleware())
	user.GET("/profile", api.AuthHandler.Profile)
	user.PUT("/profile/update", api.AuthHandler.UpdateProfile)
}
