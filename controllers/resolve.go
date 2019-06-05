package controllers

import (
	"net/http"

	"github.com/labstack/echo"
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
	IsPreferredXRPLAccountAddress string `json:"isPreferredXRPLAccountAddress"`
	Slug                          string `json:"slug"`
	PublicPage                    string `json:"publicPage"`
}

// Resolve a username of a network to an XRPL Account
func Resolve(c echo.Context) error {
	network := c.Param("network")
	username := c.Param("username")

	r := new(resolveResponse)
	// TODO: remove params here
	r.Label = "param check: " + network + " " + username
	w := new(resolveResponseWrap)
	w.Data = *r
	return c.JSON(http.StatusOK, w)
}
