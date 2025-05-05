package handlers

import (
	"fmt"
	"net/http"
	"sample/models"
	"sample/repositories"

	"github.com/labstack/echo/v4"
)

func ShowSignup(c echo.Context) error {
	errorMessage := c.QueryParam("error")
	return c.Render(http.StatusOK, "signup.html", map[string]interface{}{
		"Error": errorMessage,
	})
}
func Signup(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.Redirect(http.StatusSeeOther, "/signup?error=Missing+fields")

	}
	existingUser, err := repositories.GetUserByUsername(username)
	if err == nil && existingUser != nil {
		return c.Redirect(http.StatusSeeOther, "/signup?error=Username+already+exists")
	}
	user := models.User{
		Username: username,
		Password: password,
	}
	err = repositories.InsertUser(user)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/signup?error=Internal+server+error")
	}

	fmt.Printf("%+v\n", user)

	return c.Redirect(http.StatusSeeOther, "/login")

}
