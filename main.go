package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", hello)

	e.POST("/test", testEndpoint)

	log.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "welcome to mini server land")
}

func testEndpoint(c echo.Context) error {
	binding := make(map[string]interface{})
	bodyDecoder := json.NewDecoder(c.Request().Body)

	if err := bodyDecoder.Decode(&binding); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("%+v", binding)

	return c.JSON(http.StatusOK, binding)
}
