package account

import (
	// golang package
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var (
	errEmailEmpty    = errors.New("email is empty")
	errPasswordEmpty = errors.New("password is empty")
	errUserExist     = errors.New("user already exist!")
)

// HandleUserLogIn handles user login process.
func (h *Handler) HandleUserLogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var result userLogInResponse
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" {
		result.Code = http.StatusBadRequest
		result.Error = errEmailEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if password == "" {
		result.Code = http.StatusBadRequest
		result.Error = errPasswordEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	token, err := h.account.LogIn(context.Background(), email, password)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	result.Code = http.StatusOK
	result.Token = token
	json.NewEncoder(w).Encode(result)
}

// HandlerUserLogOut handles user logout process.
func (h *Handler) HandlerUserLogOut(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.FormValue("user_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode("failed to log out!")
		return
	}

	err = h.account.LogOut(context.Background(), userID)
	if err != nil {
		json.NewEncoder(w).Encode("failed to log out!")
		return
	}

	json.NewEncoder(w).Encode("success!")
}

// HandleUserSignUp handles user creation process.
func (h *Handler) HandleUserSignUp(w http.ResponseWriter, r *http.Request) {
	result := userSignUpResponse{
		Message: "failed!",
	}
	w.Header().Set("Content-Type", "application/json")

	bytes, err := h.infra.ReadAll(r.Body)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	var request userSignUpParam
	err = h.infra.JsonUnmarshal(bytes, &request)
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Email == "" {
		result.Code = http.StatusBadRequest
		result.Error = errEmailEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	if request.Password == "" {
		result.Code = http.StatusBadRequest
		result.Error = errPasswordEmpty.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	err = h.account.UserSignUp(context.Background(), request.Email, request.Password)
	if err != nil {
		if err == errUserExist {
			result.Code = http.StatusBadRequest
			result.Error = err.Error()

			json.NewEncoder(w).Encode(result)
			return
		}

		result.Code = http.StatusInternalServerError
		result.Error = err.Error()

		json.NewEncoder(w).Encode(result)
		return
	}

	result.Message = "success!"
	result.Code = http.StatusCreated
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) ABC(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ABC")

}
