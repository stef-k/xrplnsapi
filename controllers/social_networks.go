package controllers

import (
	"net/http"

	"github.com/stef-k/xrplnsapi/models"

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
	// usage: curl -H "XRPLNS-KEY: key-goes-value-here" http://127.0.0.1:2222/v1/social-networks
	if !models.CheckAPIKey(c.Request().Header["Xrplns-Key"]) {
		var empty interface{}
		return c.JSON(http.StatusUnauthorized, empty)
	}
	return c.JSON(http.StatusOK, networks)
}
