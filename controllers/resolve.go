package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/stef-k/xrplnsapi/models"
)

// ResponseWrap wrapper struct to provide the JSON data field
// {data: ""}
type resolveResponseWrap struct {
	Data resolveResponse `json:"data"`
}

type resolveResponse struct {
	XrplAccount                   string `json:"xrplAccount"`
	DestinationTag                string `json:"destinationTag"`
	Label                         string `json:"label"`
	IsPreferredXRPLAccountAddress bool   `json:"isPreferredXRPLAccountAddress"`
	Slug                          string `json:"slug"`
	PublicPage                    string `json:"publicPage"`
}

// Resolve a username of a network to an XRPL Account
func Resolve(c echo.Context) error {
	network := c.Param("network")
	username := c.Param("username")

	social, exists := models.GetXrplAccount(network, username)

	r := new(resolveResponse)

	if exists {
		r.XrplAccount = social.PreferredXrplAccount.XrplAccount
		r.DestinationTag = social.PreferredXrplAccount.Tag
		r.Label = social.PreferredXrplAccount.Label
		r.IsPreferredXRPLAccountAddress = social.PreferredXrplAccount.IsPreferredAddressOfUser
		r.Slug = social.User.Slug
		r.PublicPage = fmt.Sprintf("%s%s", os.Getenv("PUBLIC_USERS_URL"), social.User.Slug)
		w := new(resolveResponseWrap)
		w.Data = *r
		return c.JSON(http.StatusOK, w)
	}

	var empty interface{}
	return c.JSON(http.StatusNotFound, empty)
}
