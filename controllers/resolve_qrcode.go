package controllers

import (
	"encoding/base64"
	"net/http"

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

	var png []byte
	png, _ = qrcode.Encode(network+" "+username, qrcode.Highest, 256)
	b64 := base64.StdEncoding.EncodeToString(png)

	r := new(resolveQrcodeResponse)
	// TODO: remove params here
	r.QRCode = b64
	w := new(resolveQrcodeResponseWrap)
	w.Data = *r
	return c.JSON(http.StatusOK, w)
}
