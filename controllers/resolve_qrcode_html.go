package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/stef-k/xrplnsapi/models"
)

// resolveQrcodeResponseWrap wrapper struct to provide the JSON data field
// {data: ""}
type resolveQrcodeHTMLResponseWrap struct {
	Data resolveQrcodeHTMLResponse `json:"data"`
}

type resolveQrcodeHTMLResponse struct {
	QRCode string `json:"qrcode"`
}

// ResolveQRCodeHTML a username of a network to an XRPL Account
func ResolveQRCodeHTML(c echo.Context) error {
	// html := `<img src='data:image/png;base64,` + b64 + `' class='xrplnsQRCodeImgTag' />`
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
		html := `<img src='data:image/png;base64,` + b64 + `' class='xrplnsQRCodeImgTag' />`
		r := new(resolveQrcodeResponse)
		r.QRCode = html
		w := new(resolveQrcodeResponseWrap)
		w.Data = *r
		return c.JSON(http.StatusOK, w)
	}
	var empty interface{}
	return c.JSON(http.StatusNotFound, empty)
}
