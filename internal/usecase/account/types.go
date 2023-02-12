package account

// -------------------
// | Response Struct |
// -------------------

// JWT holds token needed for authorization.
type JWT struct {
	Token string `json:"token"`
}

// --------------------
// | Parameter Struct |
// --------------------

// UpdateUserAccountParam represents parameter needed to update an account.
type UpdateUserAccountParam struct {
	FirstName    string
	LastName     string
	RecordPeriod int
	UserID       int64
}

// UpdatePasswordParam represents parameter needed to update user's password.
type UpdatePasswordParam struct {
	Email       string
	OldPassword string
	Password    string
	UserID      int64
}
