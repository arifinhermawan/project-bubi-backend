package account

// -------------------------
// | structs for parameter |
// -------------------------

// userSignUpParam represents parameters needed to create a new user sign up.
type userSignUpParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ------------------------
// | structs for response |
// ------------------------

// userSignUpResponse represents response that will be given by endpoint /account/signup
type userSignUpResponse struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}