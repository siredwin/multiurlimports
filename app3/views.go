package app3

import(
	"github.com/labstack/echo"
	"net/http"
)


func view5(c echo.Context) (err error) {
	return c.String(http.StatusOK, "This is app3 view 5")
}

func view6(c echo.Context) (err error) {
	return c.String(http.StatusOK, "This is app3 view 6")
}