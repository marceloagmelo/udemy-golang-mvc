package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/marceloagmelo/pongor-echo"
	r "github.com/marceloagmelo/udemy-golang-mvc/routers"
)

func main() {
	e := r.App

	p := pongor.GetRenderer()
	p.Directory = "views"

	e.Renderer = p

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":3000"))

	//r.App.Start(":3000")
}
