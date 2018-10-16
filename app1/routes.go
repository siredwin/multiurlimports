package app1

import (
	"github.com/siredwin/multiurlimports/supercontext"
)

func init(){
	//Initiate the supercontext
	e := supercontext.E
	// Name the routes for app1
	app1routes := e.Group("/app1")
	app1routes.Match([]string{"GET", "POST"}, "/view1", view1)
	app1routes.Match([]string{"GET", "POST"}, "/view2", view2)

}