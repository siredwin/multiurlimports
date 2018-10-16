package app3

import (
	"github.com/siredwin/multiurlimports/supercontext"
)

func init(){
	//Initiate the supercontext
	e := supercontext.E
	// Name the routes for app3
	app1routes := e.Group("/app3")
	app1routes.Match([]string{"GET", "POST"}, "/view5", view5)
	app1routes.Match([]string{"GET", "POST"}, "/view6", view6)

}