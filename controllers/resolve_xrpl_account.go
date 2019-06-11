package controllers

import (
	"net/http"

	"github.com/stef-k/xrplnsapi/models"

	"github.com/labstack/echo"
)

type resolveXRPLAccountResponseWrap struct {
	Data []models.ResolveXRPLAccountResponse `json:"data"`
}

// ResolveXRPLAccount resolve an XRPL Account to it's Associated Users
func ResolveXRPLAccount(c echo.Context) error {
	xrplaccount := c.Param("xrplaccount")
	tag := c.Param("tag")

	response := new(resolveXRPLAccountResponseWrap)

	response.Data = models.GetAccountUsers(xrplaccount, tag)

	if len(response.Data) > 0 {
		return c.JSON(http.StatusOK, response)
	}

	var empty interface{}
	return c.JSON(http.StatusNotFound, empty)
}
