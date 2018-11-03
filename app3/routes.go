package app3

import (
	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	// Name the routes for app3
	app3routes := e.Group("/app3")
	app3routes.Match([]string{"GET", "POST"}, "/view5", view5)
	app3routes.Match([]string{"GET", "POST"}, "/view6", view6)

}
