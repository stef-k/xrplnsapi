package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

var networks = []string{
	"website",
	"email",
	"reddit",
	"twitter",
	"facebook",
	"stackexchange",
	"instagram",
	"github",
	"twitch",
	"slack",
	"discord",
	"disqus",
	"deviantart",
}

// SocialNetworks returns a list of supported social networks
func SocialNetworks(c echo.Context) error {
	return c.JSON(http.StatusOK, networks)
}
