package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/test", testEndpoint)

	log.Fatal(e.Start(":3000"))
}

func testEndpoint(c echo.Context) error {
	return c.String(200, "yo")
}
