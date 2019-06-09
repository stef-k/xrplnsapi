package models

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
