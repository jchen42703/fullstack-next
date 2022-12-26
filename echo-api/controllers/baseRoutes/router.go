package baseRoutes

import (
	"fmt"
	"net/http"

	"github.com/jchen42703/echo-api/internal/templates"
	"github.com/jchen42703/echo-api/middleware"
	"github.com/labstack/echo/v4"
	// "github.com/stock-exe/api/db"
)

func GetSession() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Print("Session: ", c.Get(middleware.ORY_SESSION_CONTEXT_KEY))
		session, ok := c.Get(middleware.ORY_SESSION_CONTEXT_KEY).(map[string]interface{})
		if !ok || session == nil {
			return c.JSON(http.StatusUnauthorized, templates.NewError(fmt.Errorf("must be signed in")))
		}

		return c.JSON(http.StatusOK, session)
	}
}

// Test routes to see session data.
func RegisterRoutes(g *echo.Group) {
	g.GET("/session", GetSession())
}
