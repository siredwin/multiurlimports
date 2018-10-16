package main

import (
	_ "github.com/siredwin/multiurlimports/app1" // app1 views from init
	_ "github.com/siredwin/multiurlimports/app2" // app2 views from init
	_ "github.com/siredwin/multiurlimports/app3" // app3 views from init
	"github.com/labstack/echo/middleware"
	"github.com/siredwin/multiurlimports/supercontext"
)


func main() {
	e := supercontext.E // use our supercontext
	e.Debug = true
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Static("/static", "static")
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
