package entities

// User represents a user returned by user.query.
type User struct {
	PHID     string   `json:"phid"`
	UserName string   `json:"userName"`
	RealName string   `json:"realName"`
	Image    string   `json:"image"`
	URI      string   `json:"uri"`
	Roles    []string `json:"roles"`
	// TM extension: this is not returned by normal Phabricator installs.
	// See https://github.com/thought-machine/phabricator for more info.
	PrimaryEmail string `json:"primaryEmail"`
}
