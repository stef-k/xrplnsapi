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
	userModel, exists := models.GetUserBySlug(slug)

	if !exists {
		var empty interface{}
		return c.JSON(http.StatusNotFound, empty)
	}

	emails = nil
	xrplAccounts = nil
	socialAccounts = nil

	// attach xrpl account in response
	if len(userModel.XrplAccounts) > 0 {
		for _, xrpl := range userModel.XrplAccounts {
			xrplAcc := new(xrplAccount)
			xrplAcc.XrplAccount = xrpl.XrplAccount
			xrplAcc.DestinationTag = xrpl.Tag
			xrplAcc.Label = xrpl.Label
			xrplAcc.IsPreferredXRPLAccountAddress = xrpl.IsPreferredAddressOfUser
			xrplAccounts = append(xrplAccounts, *xrplAcc)
		}
	}

	// attach social accounts in response
	if len(userModel.SocialAccounts) > 0 {
		for _, acc := range userModel.SocialAccounts {
			if acc.Public && acc.Verified {
				social := new(socialAccount)
				social.SocialNetwork = acc.SocialNetwork
				social.PreferredXrplAccount = acc.PreferredXrplAccount.XrplAccount
				social.DestinationTag = acc.PreferredXrplAccount.Tag
				social.Label = acc.PreferredXrplAccount.Label
				socialAccounts = append(socialAccounts, *social)
				// if social account type of email attach the sha1 version to the contact array
				if acc.SocialNetwork == "email" {
					emails = append(emails, acc.ContactMail)
				}
			}
		}
	}

	cont := new(contact)
	cont.Mail = emails

	user := new(resolveUser)

	user.Name = userModel.Name
	if userModel.PublicProfile {
		user.PublicPage = fmt.Sprintf("%s%s", os.Getenv("PUBLIC_USERS_URL"), userModel.Slug)
	}
	user.ImageURL = userModel.ImageURL
	user.SocialAccounts = socialAccounts
	user.XRPLAccounts = xrplAccounts
	user.Contact = *cont

	resp := new(resolveUserResponseWrap)
	resp.Data = *user

	return c.JSON(http.StatusOK, resp)
}
