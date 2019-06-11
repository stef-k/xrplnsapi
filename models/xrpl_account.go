package models

import (
	"fmt"
	"os"
)

// XrplAccount  model
type XrplAccount struct {
	ID                       uint64 `orm:"pk;column(id)"`
	Label                    string `orm:"column(label)"`
	XrplAccount              string `orm:"column(xrpl_account)"`
	Tag                      string `orm:"column(tag)"`
	Verified                 bool   `orm:"column(verified)"`
	IsPreferredAddressOfUser bool   `orm:"column(is_preferred_address_of_user)"`
	// Relations
	User           *User            `orm:"rel(fk);column(user_id)"`
	SocialAccounts []*SocialAccount `orm:"reverse(many)"`
}

// TableName custom table name
func (x *XrplAccount) TableName() string {
	return "xrpl_accounts"
}

// ResolveXRPLAccountResponse response container
type ResolveXRPLAccountResponse struct {
	Name                          string `json:"name"`
	Slug                          string `json:"slug"`
	PublicPage                    string `json:"publicPage"`
	IsPreferredXRPLAccountAddress bool   `json:"isPreferredXRPLAccountAddress"`
}

// GetAccountUsers get all users for this XRPL account
func GetAccountUsers(xrplaccount, tag string) []ResolveXRPLAccountResponse {
	var xrplAccounts []ResolveXRPLAccountResponse
	var accounts []XrplAccount
	var account XrplAccount

	if len(xrplaccount) <= 24 {
		return xrplAccounts
	}

	qs := DB.QueryTable("xrpl_accounts")

	if len(tag) > 0 {
		qs.Filter("xrpl_account", xrplaccount).Filter("tag", tag).All(&accounts)
		if len(accounts) > 0 {
			for _, account = range accounts {
				DB.LoadRelated(&account, "User")
				if !account.User.Locked && account.User.PublicProfile {
					var response ResolveXRPLAccountResponse
					response.Name = account.User.Name
					response.Slug = account.User.Slug
					response.IsPreferredXRPLAccountAddress = account.IsPreferredAddressOfUser
					response.PublicPage = fmt.Sprintf("%s%s", os.Getenv("PUBLIC_USERS_URL"), account.User.Slug)
					xrplAccounts = append(xrplAccounts, response)
				}
			}
		}
	} else {
		qs.Filter("xrpl_account", xrplaccount).Filter("tag__isnull", true).All(&accounts)
		if len(accounts) > 0 {
			for _, account = range accounts {
				DB.LoadRelated(&account, "User")
				if !account.User.Locked && account.User.PublicProfile {
					var response ResolveXRPLAccountResponse
					response.Name = account.User.Name
					response.Slug = account.User.Slug
					response.IsPreferredXRPLAccountAddress = account.IsPreferredAddressOfUser
					response.PublicPage = fmt.Sprintf("%s%s", os.Getenv("PUBLIC_USERS_URL"), account.User.Slug)
					xrplAccounts = append(xrplAccounts, response)
				}
			}
		}
	}

	return xrplAccounts
}
