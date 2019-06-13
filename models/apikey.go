package models

// APIKey model
type APIKey struct {
	ID     uint64 `orm:"pk;column(id)"`
	APIKey string `orm:"column(api_key)"`
	Locked bool   `orm:"column(locked)"`
	// Relations
	User *User `orm:"rel(fk);column(user_id)"`
}

// TableName custom table name
func (a *APIKey) TableName() string {
	return "api_keys"
}

// CheckAPIKey checks the given API key
func CheckAPIKey(key []string) bool {
	// check and clean the incoming key from HTTP headers
	var theKey string
	if len(key) > 0 {
		theKey = key[0]
	} else {
		return false
	}
	var apiKey APIKey
	qs := DB.QueryTable("api_keys")
	exists := qs.Filter("api_key", theKey).Exist()

	if exists {
		qs.Filter("api_key", theKey).One(&apiKey)
		DB.LoadRelated(&apiKey, "User")
		return (!apiKey.Locked && !apiKey.User.Locked)
	}

	return false
}
