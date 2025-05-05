package main

import (
	"html/template"
	"io"
	"sample/handlers"
	"sample/storage"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.GET("/", handlers.Home)
	storage.InitDB()

	e.GET("/signup", handlers.ShowSignup)
	e.POST("/signup", handlers.Signup)
	e.GET("/login", handlers.ShowLogin)
	e.POST("/login", handlers.ValidateUser)
	e.GET("/:username", handlers.WelcomePage)

	e.Logger.Fatal(e.Start(":8080"))
}
