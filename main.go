package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Get("/", hello)

	e.Logger().Fatal(e.Start("3000"))
}
