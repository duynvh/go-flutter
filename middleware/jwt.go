package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-flutter/model"
	"go-flutter/security"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.SECRET_KEY),
	}

	return middleware.JWTWithConfig(config)
}
