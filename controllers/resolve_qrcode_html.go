package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo"
	qrcode "github.com/skip2/go-qrcode"
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
	network := c.Param("network")
	username := c.Param("username")

	var png []byte
	png, _ = qrcode.Encode(network+" "+username, qrcode.Highest, 256)
	b64 := base64.StdEncoding.EncodeToString(png)
	html := `<img src='data:image/png;base64,` + b64 + `' class='xrplnsQRCodeImgTag' />`

	r := new(resolveQrcodeHTMLResponse)
	// TODO: remove params here
	r.QRCode = html
	w := new(resolveQrcodeHTMLResponseWrap)
	w.Data = *r
	return c.JSON(http.StatusOK, w)
}
