package handlers

import (
	"net/http"
	"sample/repositories"

	"github.com/labstack/echo/v4"
)

func ShowLogin(c echo.Context) error {
	errorMessage := c.QueryParam("error")
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"Error": errorMessage,
	})
}
func ValidateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.Redirect(http.StatusSeeOther, "/login?error=Username+and+password+are+required")
	}

	// Fetch user by username
	_, err := repositories.GetUserByUsername(username)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login?error=Invalid+username+or+password")
	}

	return c.Redirect(http.StatusSeeOther, "/"+username)

}
