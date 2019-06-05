package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// ResponseWrap wrapper struct to provide the JSON data field
// {data: ""}
type resolveUserResponseWrap struct {
	Data resolveUser `json:"data"`
}

type resolveUser struct {
	Name           string          `json:"name"`
	PublicPage     string          `json:"publicPage"`
	ImageURL       string          `json:"imageUrl"`
	Contact        contact         `json:"contact"`
	SocialAccounts []socialAccount `json:"socialAccounts"`
	XRPLAccounts   []xrplAccount   `json:"xrplAccounts"`
}

type contact struct {
	Mail []string `json:"mail"`
}

type socialAccount struct {
	SocialNetwork        string `json:"socialNetwork"`
	PreferredXrplAccount string `json:"preferredXrplAccount"`
	Label                string `json:"label"`
	DestinationTag       string `json:"destinationTag"`
}

type xrplAccount struct {
	XrplAccount                   string `json:"xrplAccount"`
	DestinationTag                string `json:"destinationTag"`
	Label                         string `json:"label"`
	IsPreferredXRPLAccountAddress bool   `json:"isPreferredXRPLAccountAddress"`
}

var xrplAccounts = []xrplAccount{}
var socialAccounts = []socialAccount{}
var emails []string

// ResolveUser resolves a User account based on a slug
func ResolveUser(c echo.Context) error {
	slug := c.Param("slug")

	emails = append(emails, "one@example.com")
	emails = append(emails, "two@example.com")
	cont := new(contact)
	cont.Mail = emails

	xrplAcc := new(xrplAccount)
	xrplAcc.Label = "XRPL Account 1"

	xrplAccounts = append(xrplAccounts, *xrplAcc)

	xrplAcc = new(xrplAccount)
	xrplAcc.Label = "XRPL Account 2"

	xrplAccounts = append(xrplAccounts, *xrplAcc)

	socialAcc := new(socialAccount)
	socialAcc.Label = "Social Account 1"

	socialAccounts = append(socialAccounts, *socialAcc)

	socialAcc = new(socialAccount)
	socialAcc.Label = "Social Account 2"

	socialAccounts = append(socialAccounts, *socialAcc)

	user := new(resolveUser)

	user.Name = "check param slug: " + slug
	user.SocialAccounts = socialAccounts
	user.XRPLAccounts = xrplAccounts
	user.Contact = *cont

	resp := new(resolveUserResponseWrap)
	resp.Data = *user

	return c.JSON(http.StatusOK, resp)
}
