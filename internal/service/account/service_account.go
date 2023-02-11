package account

import (
	// golang package
	"context"
	"errors"
	"log"

	// external package
	"golang.org/x/crypto/bcrypt"
)

var (
	errIncorrectPassword = errors.New("incorrect password!")

	// for mocking purpose
	compareHashPassword  = bcrypt.CompareHashAndPassword
	generateFromPassword = bcrypt.GenerateFromPassword
)

// CheckPasswordCorrect will check whether user's password match with current password or not.
func (svc *Service) CheckPasswordCorrect(ctx context.Context, email, password string) error {
	meta := map[string]interface{}{
		"email": email,
	}

	account, err := svc.rsc.GetUserAccountByEmailFromDB(ctx, email)
	if err != nil {
		log.Printf("[CheckPasswordCorrect] svc.rsc.GetUserAccountByEmailFromDB() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	err = compareHashPassword([]byte(account.Password), []byte(password))
	if err != nil {
		log.Printf("[CheckPasswordCorrect] compareHashPassword() got an error: %+v\nMeta:%+v\n", errIncorrectPassword, meta)
		return errIncorrectPassword
	}

	return nil
}

// GetUserAccountByEmail will check whether an account is already exist by using email.
func (svc *Service) GetUserAccountByEmail(ctx context.Context, email string) (Account, error) {
	account, err := svc.rsc.GetUserAccountByEmailFromDB(ctx, email)
	if err != nil {
		meta := map[string]interface{}{
			"email": email,
		}

		log.Printf("[GetUserAccountByEmail] svc.rsc.GetUserAccountByEmailFromDB() got an error: %+v\nMeta:%+v\n", err, meta)
		return Account{}, err
	}

	return Account(account), nil
}

// InsertUserAccount will create a new user account.
func (svc *Service) InsertUserAccount(ctx context.Context, email, password string) (err error) {
	meta := map[string]interface{}{
		"email": email,
	}

	hashed, err := generateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[InsertUserAccount] generateFromPassword() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	err = svc.rsc.InsertUserAccountToDB(ctx, email, string(hashed))
	if err != nil {
		log.Printf("[InsertUserAccount] svc.rsc.InsertUserAccountToDB() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}

// UpdateUserAccount will update the information of an existing user account.
func (svc *Service) UpdateUserAccount(ctx context.Context, param UpdateUserAccountParam) error {
	meta := map[string]interface{}{
		"first_name":    param.FirstName,
		"last_name":     param.LastName,
		"record_period": param.RecordPeriod,
		"user_id":       param.UserID,
	}

	err := svc.rsc.UpdateUserAccountInDB(ctx, param)
	if err != nil {
		log.Printf("[UpdateUserAccount] svc.rsc.UpdateUserAccountInDB() got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	return nil
}
