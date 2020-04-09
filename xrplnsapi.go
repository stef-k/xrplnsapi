package main

import (
	"net/http"

	"github.com/stef-k/xrplnsapi/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi, please check https://xrplns.com/docs")
	})
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start("127.0.0.1:2222"))
}
