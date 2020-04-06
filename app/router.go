package main

import (
	"wego/app/controllers"
	"wego/library"
)

var Routers = []library.Router{
	library.NewRouter("GET", "/", controllers.GreetingsHandler),
}
