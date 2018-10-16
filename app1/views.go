package app1

import(
	"github.com/labstack/echo"
	"net/http"
)


func view1(c echo.Context) (err error) {
	return c.String(http.StatusOK, "This is app1 view 1")
}

func view2(c echo.Context) (err error) {
	return c.String(http.StatusOK, "This is app1 view 2")
}