package middleware

import (
	"net/http"

	"github.com/jchen42703/echo-api/auth"
	"github.com/jchen42703/echo-api/internal/templates"
	echo "github.com/labstack/echo/v4"
	ory "github.com/ory/client-go"
)

const ORY_SESSION_CONTEXT_KEY = "user_session"

// Checks if a user has an active session (is logged in) by checking the session cookie.
// Then, it adds the session object to the "user_session" key in the echo Context.
func CreateAuthMiddleware(oryClient *ory.APIClient) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			whitelist := []string{
				"/api",
				"/api/notprotected", // dev route for not protected route
				"/api/auth/signup",
				"/api/auth/logout",
			}

			// Skip auth if whitelisted
			for _, whitelistedPath := range whitelist {
				if whitelistedPath == c.Path() {
					return next(c)
				}
			}

			// check if we have a session
			session, _, err := auth.ValidateSession(oryClient, c)
			if err != nil {
				// fmt.Println("validate sess err: ", templates.NewError(err))
				// return c.JSON(http.StatusUnauthorized, templates.NewError(err))
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
			}

			isLogin := c.Path() == "/api/auth/login"

			// // Force validate the route behind authentication
			// // Checks for existing JWT "session"
			// jwtPublicSecret := []byte(os.Getenv("JWT_PUBLIC_KEY"))
			// jwtName := os.Getenv("JWT_COOKIE_NAME")
			// // Check for valid login JWT cookie
			// jwtFromCookie, err := c.Cookie(jwtName)
			// if err != nil {
			// 	if isLogin {
			// 		return next(c)
			// 	}

			// 	return c.JSON(http.StatusUnauthorized, templates.NewError(fmt.Errorf("must be logged in to use this endpoint")))
			// }

			// // TODO: refresh with new jwt if valid JWT
			if isLogin {
				return c.JSON(http.StatusOK, templates.NewResp("already logged in"))
			}

			// so that it can be reused in next handlers
			c.Set(ORY_SESSION_CONTEXT_KEY, session)
			return next(c)
		}
	}
}
