package app2

import (
	"github.com/siredwin/multiurlimports/supercontext"
)

func init(){
	//Initiate the supercontext
	e := supercontext.E
	// Name the routes for app2
	app2routes := e.Group("/app2")
	app2routes.Match([]string{"GET", "POST"}, "/view3", view3)
	app2routes.Match([]string{"GET", "POST"}, "/view4", view4)

}
