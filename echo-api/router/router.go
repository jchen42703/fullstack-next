package router

import (
	customMiddleware "github.com/jchen42703/echo-api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	ory "github.com/ory/client-go"
)

func New(oryClient *ory.APIClient) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4455", "http://localhost:4455/"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true, // This is necessary if you need session cookies because fetch with credentials:true requires CORS to return the credentials header as true
	}))
	e.Use(middleware.CSRF())
	e.Use(customMiddleware.CreateAuthMiddleware(oryClient))
	e.Validator = NewValidator()
	return e
}
