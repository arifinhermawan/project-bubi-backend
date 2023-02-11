package account

// -------------------------
// | structs for parameter |
// -------------------------

// userLogInParam represents parameters needed to create a new user sign up.
type userLogInParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// userSignUpParam represents parameters needed to create a new user sign up.
type userSignUpParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// updateUserAccount represents parameters needed to update user account.
type updateUserAccount struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	RecordPeriod int    `json:"record_period"`
	UserID       int64  `json:"user_id"`
}

// ------------------------
// | structs for response |
// ------------------------

// defaultResponse represents default response of an API call
type defaultResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// userSignUpResponse represents response that will be given by endpoint /account/signup
type userLogInResponse struct {
	defaultResponse
	Token string `json:"token"`
}

// userSignUpResponse represents response that will be given by endpoint /account/signup
type userSignUpResponse struct {
	defaultResponse
	Message string `json:"message"`
}
