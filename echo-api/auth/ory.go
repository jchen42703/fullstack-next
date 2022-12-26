package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	ory "github.com/ory/client-go"
)

func NewOryAPIClient(baseKratosUrl string) *ory.APIClient {
	if baseKratosUrl == "" {
		baseKratosUrl = "http://localhost:4433"
	}

	// register a new Ory client with the URL set to the Ory CLI Proxy
	// we can also read the URL from the env or a config file
	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("%s/.ory", baseKratosUrl)}}
	return ory.NewAPIClient(c)
}

func ValidateSession(client *ory.APIClient, ctx echo.Context) (*ory.Session, *http.Response, error) {
	fmt.Println("Cookies: ", ctx.Cookies())
	authCookie, err := ctx.Cookie("ory_kratos_session")
	if err != nil {
		return nil, nil, fmt.Errorf("ValidateSession: could not find session cookie 'ory_kratos_session'")
	}

	fmt.Println(authCookie.String())
	session, resp, err := client.FrontendApi.ToSession(ctx.Request().Context()).Cookie(authCookie.String()).Execute()
	if (err != nil && session == nil) || (err == nil && !*session.Active) {
		// Not logged in
		return nil, nil, fmt.Errorf("ValidateSession: no valid session: %s", err)
	}

	return session, resp, nil
}
