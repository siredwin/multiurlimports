package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/siredwin/multiurlimports/app1"
	"github.com/siredwin/multiurlimports/app2"
	"github.com/siredwin/multiurlimports/app3"
)

func main() {
	e := echo.New()
	e.Debug = true
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Static("/static", "static")

	// Register apps
	app1.New(e)
	app2.New(e)
	app3.New(e)

	// #*************** START SERVER *****************# //
	e.Logger.Fatal(e.Start(":8000"))
}

//Try these routes:
//http://localhost:8000/app1/view1
//http://localhost:8000/app1/view2
//http://localhost:8000/app2/view3
//http://localhost:8000/app2/view4
//http://localhost:8000/app3/view5
//http://localhost:8000/app3/view6
