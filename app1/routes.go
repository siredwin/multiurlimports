package app1

import (
	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	// Name the routes for app1
	app1routes := e.Group("/app1")
	app1routes.Match([]string{"GET", "POST"}, "/view1", view1)
	app1routes.Match([]string{"GET", "POST"}, "/view2", view2)

}
