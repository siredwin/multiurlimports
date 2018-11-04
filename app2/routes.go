package app2

import (
	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	// Name the routes for app2
	app2routes := e.Group("/app2")
	app2routes.Match([]string{"GET", "POST"}, "/view3", view3)
	app2routes.Match([]string{"GET", "POST"}, "/view4", view4)

}
