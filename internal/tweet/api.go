package tweet

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Group, service Service, handleScopes func(...string) echo.MiddlewareFunc) {
	issue := tweet{service}
	router := r.Group("/tweets")
	{
		router.GET("", issue.Query, handleScopes("tweet.read", "users.read"))
		router.GET("/:id", issue.Get, handleScopes("tweet.read", "users.read"))
		router.POST("", issue.Create, handleScopes("tweet.read", "tweet.write", "users.read"))
		router.DELETE("/:id", issue.Delete, handleScopes("tweet.read", "tweet.write", "users.read"))
		// router.GET("", issue.Query, handleScopes())
		// router.GET("/:id", issue.Get, handleScopes())
		// router.POST("", issue.Create, handleScopes())
		// router.DELETE("/:id", issue.Delete, handleScopes())
	}
}

type tweet struct {
	service Service
}

func (r tweet) Get(c echo.Context) error {

	tweet, err := r.service.Get(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, tweet)
}

func (r tweet) Delete(c echo.Context) error {

	tweet, err := r.service.Get(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	username := getUsername(c)

	if tweet.User != username {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	err = r.service.Delete(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}

func (r tweet) Create(c echo.Context) error {

	username := getUsername(c)

	var input CreateTweetRequest
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs. Please check your inputs")
	}
	input.User = username
	issue, err := r.service.Create(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return c.JSON(http.StatusCreated, issue)
}

func (r tweet) Query(c echo.Context) error {

	issue, err := r.service.Query(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, issue)
}

func getUsername(c echo.Context) string {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return username
}
