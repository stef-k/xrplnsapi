package models

// SocialAccount model
type SocialAccount struct {
	ID            uint64 `orm:"pk;column(id)"`
	SocialNetwork string `orm:"column(social_network)"`
	Username      string `orm:"column(username)"`
	ContactMail   string `orm:"column(contact_mail)"`
	Verified      bool   `orm:"column(verified)"`
	Public        bool   `orm:"column(public)"`
	// Foreign Key to User
	User                 *User        `orm:"rel(fk);column(user_id)"`
	PreferredXrplAccount *XrplAccount `orm:"rel(fk);column(preferred_xrpl_account)"`
}

// TableName custom table name
func (s *SocialAccount) TableName() string {
	return "social_accounts"
}

// GetXrplAccount return the preferred XRPL Account by searching a social network for a username
func GetXrplAccount(network, username string) (socialAccount SocialAccount, exists bool) {
	var social SocialAccount
	maybeSha := false

	qs := DB.QueryTable("social_accounts")
	exists = qs.Filter("social_network", network).Filter("username", username).Filter("public", true).Filter("verified", true).Exist()
	if !exists && network == "email" {
		exists = qs.Filter("contact_mail", username).Filter("public", true).Filter("verified", true).Exist()
		maybeSha = true
	}
	if exists {
		if !maybeSha {
			qs.Filter("social_network", network).Filter("username", username).Filter("public", true).Filter("verified", true).One(&social)
		} else {
			// check if someone searches using sha1 version of mail
			qs.Filter("contact_mail", username).Filter("public", true).Filter("verified", true).One(&social)
		}
		DB.LoadRelated(&social, "PreferredXrplAccount")
		DB.LoadRelated(&social, "User")
		if !social.User.Locked {
			return social, exists
		}
	}
	return social, false
}
