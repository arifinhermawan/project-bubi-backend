package account

import (
	// golang package
	"context"
	"errors"
	"log"

	// internal package
	"github.com/arifinhermawan/bubi/internal/service/account"
)

var (
	errUserExist    = errors.New("user already exist!")
	errUserNotExist = errors.New("user not exist!")
)

// LogIn handles the log in process for a user.
// It will check the existence of a user first.
// If it exist, then it will continue the log in process.
func (uc *UseCase) LogIn(ctx context.Context, email, password string) (string, error) {
	meta := map[string]interface{}{
		"email": email,
	}

	acc, err := uc.account.GetUserAccountByEmail(ctx, email)
	if err != nil {
		log.Printf("[LogIn] uc.account.GetUserAccountByEmail() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	accountNotExist := acc == account.Account{}
	if accountNotExist {
		log.Printf("[LogIn] User not exist!\nMeta:%+v\n", meta)
		return "", errUserNotExist
	}

	err = uc.account.CheckPasswordCorrect(ctx, email, password)
	if err != nil {
		log.Printf("[LogIn] uc.account.CheckPasswordCorrect() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	token, err := uc.account.GenerateJWT(ctx, acc.ID, email)
	if err != nil {
		log.Printf("[LogIn] uc.account.CheckIsAccountExist() got an error: %+v\nMeta:%+v\n", err, meta)
		return "", err
	}

	return token, nil
}

// LogOut handles the log out process for a user.
func (uc *UseCase) LogOut(ctx context.Context, userID int64) error {
	meta := map[string]interface{}{
		"user_id": userID,
	}

	err := uc.account.InvalidateJWT(ctx, userID)
	if err != nil {
		log.Printf("[LogOut] uc.account.InvalidateJWT() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}

// UserSignUp will process the creation of user account.
// Before creating a new account, it'll check whether that account exist or not.
// If it's a new account, then it'll create a new user account.
func (uc *UseCase) UserSignUp(ctx context.Context, email, password string) error {
	meta := map[string]interface{}{
		"email": email,
	}

	acc, err := uc.account.GetUserAccountByEmail(ctx, email)
	if err != nil {
		log.Printf("[LogIn] uc.account.GetUserAccountByEmail() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	accountExist := acc != account.Account{}
	if accountExist {
		log.Printf("[UserSignUp] User already exist!\nMeta:%+v\n", meta)
		return errUserExist
	}

	err = uc.account.InsertUserAccount(ctx, email, password)
	if err != nil {
		log.Printf("[UserSignUp] uc.account.InsertUserAccount() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}
