package app2

import(
	"github.com/labstack/echo"
	"net/http"
)


func view3(c echo.Context) (err error) {
	return c.String(http.StatusOK, "This is app2 view 3")
}

func view4(c echo.Context) (err error) {
	return c.String(http.StatusOK, "This is app2 view 4")
}