package account

import (
	// golang package
	"context"
	"errors"
	"log"
)

var (
	errUserExist = errors.New("user already exist!")
)

// UserSignUp will process the creation of user account.
// Before creating a new account, it'll check whether that account exist or not.
// If it's a new account, then it'll create a new user account.
func (uc *UseCase) UserSignUp(ctx context.Context, email, password string) error {
	meta := map[string]interface{}{
		"email": email,
	}

	isUserExist, err := uc.account.CheckIsAccountExist(ctx, email)
	if err != nil {
		log.Printf("[UserSignUp] uc.account.CheckIsAccountExist() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	if isUserExist {
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
