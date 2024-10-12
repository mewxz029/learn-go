package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mewxz029/02_package/calculator"
)

func main() {
	fmt.Println(calculator.Add(1, 2))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
