package account

import (
	// golang package
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errEmailEmpty    = errors.New("email is empty")
	errPasswordEmpty = errors.New("password is empty")
	errUserExist     = errors.New("user already exist!")
)

// HandleUserSignUp handles user creation process
func (h *Handler) HandleUserSignUp(w http.ResponseWriter, r *http.Request) {
	result := userSignUpResponse{
		Message: "failed!",
	}

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
