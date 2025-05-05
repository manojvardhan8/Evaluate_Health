package handlers

import (
	"net/http"
	"sample/repositories"

	"github.com/labstack/echo/v4"
)

func WelcomePage(c echo.Context) error {
	username := c.Param("username")
	user, err := repositories.GetUserByUsername(username)
	if err != nil || user == nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}
	return c.Render(http.StatusOK, "welcome.html", map[string]interface{}{
		"Username": username,
	})

}
