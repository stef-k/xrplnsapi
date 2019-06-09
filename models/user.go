package models

// User model
type User struct {
	ID            uint64 `orm:"pk;column(id)"`
	Name          string `orm:"column(name)"`
	PublicProfile bool   `orm:"column(public_profile)"`
	ImageURL      string `orm:"column(image_url)"`
	Slug          string `orm:"column(slug)"`
	Locked        bool   `orm:"column(account_locked)"`
	// Reverse relationships
	SocialAccounts []*SocialAccount `orm:"reverse(many)"`
	XrplAccounts   []*XrplAccount   `orm:"reverse(many)"`
}

// TableName custom table name
func (u *User) TableName() string {
	return "users"
}

// GetUserBySlug find a user by slug
func GetUserBySlug(slug string) (user User, exists bool) {
	var u User
	qs := DB.QueryTable("users")
	exists = qs.Filter("slug", slug).Filter("account_locked", false).Filter("public_profile", true).Exist()
	if exists {
		qs.Filter("slug", slug).Filter("account_locked", false).Filter("public_profile", true).One(&u)
		DB.LoadRelated(&u, "XrplAccounts")
		// for _, xrpl := range u.XrplAccounts {
		// 	DB.LoadRelated(xrpl, "")
		// }
		DB.LoadRelated(&u, "SocialAccounts")
		for _, social := range u.SocialAccounts {
			DB.LoadRelated(social, "PreferredXrplAccount")
		}

		return u, exists
	}
	return u, false
}
