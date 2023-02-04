package account

import (
	// golang package
	"context"
	"log"

	// external package
	"golang.org/x/crypto/bcrypt"

	// internal package
	"github.com/arifinhermawan/bubi/internal/entity"
)

var (
	// for mocking purpose
	compareHashPassword  = bcrypt.CompareHashAndPassword
	generateFromPassword = bcrypt.GenerateFromPassword
)

// CheckIsAccountExist will check whether an account is already exist by using email.
func (svc *Service) CheckIsAccountExist(ctx context.Context, email string) (bool, error) {
	account, err := svc.rsc.GetUserAccountByEmailFromDB(ctx, email)
	if err != nil {
		meta := map[string]interface{}{
			"email": email,
		}

		log.Printf("[CheckIsAccountExist] svc.rsc.GetUserAccountByEmailFromDB() got an error: %+v\nMeta:%+v\n", err, meta)
		return false, err
	}

	isAccountExist := account != entity.Account{}
	return isAccountExist, nil
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
