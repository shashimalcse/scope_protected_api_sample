package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"sample_api/internal/tweet"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
)

func main() {

	e := buildHandler()
	e.Logger.Fatal(e.Start(":8081"))
}

func buildHandler() *echo.Echo {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	rg := e.Group("/api")
	config := echojwt.Config{
		KeyFunc: getKey,
	}
	rg.Use(echojwt.WithConfig(config))
	tweet.RegisterHandlers(rg, tweet.NewService(), handleScopes)
	return e
}

func getKey(token *jwt.Token) (interface{}, error) {

	keySet, err := jwk.Fetch(context.Background(), "https://dev.api.asgardeo.io/t/thilinas/oauth2/jwks")
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)

	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}

	return pubkey, nil
}

func handleScopes(requiredScopes ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			scope := claims["scope"].(string)
			scopes := strings.Fields(scope)

			for _, s := range requiredScopes {
				found := false
				for _, u := range scopes {
					if s == u {
						found = true
						break
					}
				}
				if !found {
					return c.JSON(http.StatusUnauthorized, "Insufficient Scopes")
				}
			}

			return next(c)
		}
	}
}
