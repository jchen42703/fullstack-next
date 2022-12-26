package baseRoutes

import (
	"fmt"
	"net/http"

	"github.com/jchen42703/echo-api/internal/templates"
	"github.com/jchen42703/echo-api/middleware"
	"github.com/labstack/echo/v4"

	// "github.com/stock-exe/api/db"
	ory "github.com/ory/client-go"
)

func GetSession() echo.HandlerFunc {
	return func(c echo.Context) error {
		session, ok := c.Get(middleware.ORY_SESSION_CONTEXT_KEY).(*ory.Session)
		if !ok || session == nil {
			return c.JSON(http.StatusUnauthorized, templates.NewError(fmt.Errorf("must be signed in")))
		}

		return c.JSON(http.StatusOK, session)
	}
}

// Random protected endpoint
func GetProtected() echo.HandlerFunc {
	return func(c echo.Context) error {
		session, ok := c.Get(middleware.ORY_SESSION_CONTEXT_KEY).(*ory.Session)
		if !ok || session == nil {
			return c.JSON(http.StatusUnauthorized, templates.NewError(fmt.Errorf("must be signed in")))
		}

		return c.JSON(http.StatusOK, "protected")
	}
}

// Test routes to see session data.
func RegisterRoutes(g *echo.Group) {
	g.GET("/session", GetSession())
	g.GET("/protected", GetProtected())
}
