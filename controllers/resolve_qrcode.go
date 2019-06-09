package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/stef-k/xrplnsapi/models"

	"github.com/labstack/echo"
	qrcode "github.com/skip2/go-qrcode"
)

// resolveQrcodeResponseWrap wrapper struct to provide the JSON data field
// {data: ""}
type resolveQrcodeResponseWrap struct {
	Data resolveQrcodeResponse `json:"data"`
}

type resolveQrcodeResponse struct {
	QRCode string `json:"qrcode"`
}

// ResolveQRCode a username of a network to an XRPL Account
func ResolveQRCode(c echo.Context) error {
	network := c.Param("network")
	username := c.Param("username")

	account, exists := models.GetXrplAccount(network, username)

	if exists {
		var png []byte
		if len(account.PreferredXrplAccount.Tag) > 0 {
			png, _ = qrcode.Encode(account.PreferredXrplAccount.XrplAccount+"&dt="+account.PreferredXrplAccount.Tag, qrcode.Highest, 256)
		} else {
			png, _ = qrcode.Encode(account.PreferredXrplAccount.XrplAccount, qrcode.Highest, 256)
		}
		b64 := base64.StdEncoding.EncodeToString(png)

		r := new(resolveQrcodeResponse)
		r.QRCode = b64
		w := new(resolveQrcodeResponseWrap)
		w.Data = *r
		return c.JSON(http.StatusOK, w)
	}
	var empty interface{}
	return c.JSON(http.StatusNotFound, empty)
}
