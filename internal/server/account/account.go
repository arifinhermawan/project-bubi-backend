package account

import (
	// golang package
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	// golang package
	"github.com/arifinhermawan/bubi/internal/usecase/account"
)

const (
	emailKey        = "email"
	firstNameKey    = "first_name"
	lastNameKey     = "last_name"
	passwordKey     = "password"
	recordPeriodKey = "record_period"
	userIDKey       = "user_id"
)

var (
	errEmailEmpty          = errors.New("email is empty")
	errNameEmpty           = errors.New("name is empty")
	errPasswordEmpty       = errors.New("password is empty")
	errRecordPeriodInvalid = errors.New("record_period not valid")
	errUserExist           = errors.New("user already exist!")
	errUserIDInvalid       = errors.New("user_id not valid")
)

// HandleUserLogIn handles user login process.
func (h *Handler) HandleUserLogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var result userLogInResponse
	email := r.FormValue(emailKey)
	password := r.FormValue(passwordKey)

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

	token, err := h.account.LogIn(context.Background(), strings.ToLower(email), password)
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
	userIDStr := r.FormValue(userIDKey)

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

// HandleUpdateUserAccount will perform an update on user account
func (h *Handler) HandleUpdateUserAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := defaultResponse{
		Error: "",
		Code:  http.StatusBadRequest,
	}

	bytes, err := h.infra.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Code = http.StatusBadRequest
		response.Error = err.Error()

		json.NewEncoder(w).Encode(response)
		return
	}

	var request updateUserAccount

	err = h.infra.JsonUnmarshal(bytes, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Code = http.StatusBadRequest
		response.Error = err.Error()

		json.NewEncoder(w).Encode(response)
		return
	}

	if request.FirstName == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = errNameEmpty.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.RecordPeriod <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = errRecordPeriodInvalid.Error()
		json.NewEncoder(w).Encode(response)

		return
	}

	if request.UserID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = errUserIDInvalid.Error()
		json.NewEncoder(w).Encode(response)

		return
	}

	err = h.account.UpdateUserAccount(context.Background(), account.UpdateUserAccountParam{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		RecordPeriod: request.RecordPeriod,
		UserID:       request.UserID,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Code = http.StatusInternalServerError
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)

		return
	}

	w.WriteHeader(http.StatusOK)
	response.Code = http.StatusOK
	response.Error = ""
	json.NewEncoder(w).Encode(response)
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

	err = h.account.UserSignUp(context.Background(), strings.ToLower(request.Email), request.Password)
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
