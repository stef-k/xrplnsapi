package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type resolveXRPLAccountResponseWrap struct {
	Data []resolveXRPLAccountResponse `json:"data"`
}

type resolveXRPLAccountResponse struct {
	Name                          string `json:"name"`
	Slug                          string `json:"slug"`
	PublicPage                    string `json:"publicPage"`
	IsPreferredXRPLAccountAddress string `json:"isPreferredXRPLAccountAddress"`
}

// ResolveXRPLAccount resolve an XRPL Account to it's Associated Users
func ResolveXRPLAccount(c echo.Context) error {
	xrplaccount := c.Param("xrplaccount")
	tag := c.Param("tag")

	var results = []resolveXRPLAccountResponse{}

	r := new(resolveXRPLAccountResponse)
	r.Name = "param check: " + xrplaccount + " " + tag

	results = append(results, *r)

	r = new(resolveXRPLAccountResponse)
	r.Name = "second result"

	results = append(results, *r)

	w := new(resolveXRPLAccountResponseWrap)

	w.Data = results

	return c.JSON(http.StatusOK, w)
}
