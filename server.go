package main

import (
	"net/http"

	"github.com/stef-k/xrplnsapi/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start("127.0.0.1:2222"))
}
