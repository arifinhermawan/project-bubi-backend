package account

import (
	// golang package
	"context"
	"database/sql"
	"log"

	// internal package
	"github.com/arifinhermawan/bubi/internal/entity"
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
)

// GetUserAccountByEmailFromDB will fetch user's information based of account's email.
func (rsc *Resource) GetUserAccountByEmailFromDB(ctx context.Context, email string) (entity.Account, error) {
	meta := map[string]interface{}{
		"email": email,
	}

	account, err := rsc.db.GetUserAccountByEmail(ctx, email)
	if err != nil {
		log.Printf("[GetUserAccountByEmailFromDB] rsc.db.GetUserAccountByEmail() got an error: %+v\nMeta: %+v\n", err, meta)
		return entity.Account{}, err
	}

	entity := entity.Account{
		ID:                account.ID,
		FirstName:         account.FirstName.String,
		LastName:          account.LastName.String,
		Password:          account.Password,
		Email:             account.Email,
		RecordPeriodStart: account.RecordPeriodStart,
	}

	return entity, nil
}

// InsertUserAccountToDB will create a new entry of user account in database.
func (rsc *Resource) InsertUserAccountToDB(ctx context.Context, email, password string) error {
	meta := map[string]interface{}{
		"email": email,
	}

	var err error
	tx, err := rsc.db.BeginTX(ctx, nil)
	if err != nil {
		log.Printf("[InsertUserAccountToDB] rsc.db.BeginTX() got an error: %+v\nMeta: %+v\n", err, meta)
		return err
	}

	defer func() {
		err := rsc.rollbackTX(ctx, tx, err)
		if err != nil {
			log.Printf("[InsertUserAccountToDB] rsc.rollbackTX() got an error: %+v\nMeta: %+v\n", err, meta)
		}
	}()

	err = rsc.db.InsertUserAccount(ctx, tx, email, password)
	if err != nil {
		log.Printf("[InsertUserAccountToDB] rsc.db.InsertUserAccount() got an error: %+v\nMeta: %+v\n", err, meta)
		return err
	}

	errCommit := rsc.db.Commit(tx)
	if errCommit != nil {
		log.Printf("[InsertUserAccountToDB] rsc.db.Commit() got an error: %+v\nMeta: %+v\n", errCommit, meta)
	}

	return nil
}

// UpdateUserAccountInDB will update user's account based on the given parameter.
func (rsc *Resource) UpdateUserAccountInDB(ctx context.Context, param UpdateUserAccountParam) error {
	meta := map[string]interface{}{
		"first_name":    param.FirstName,
		"last_name":     param.LastName,
		"record_period": param.RecordPeriod,
		"user_id":       param.UserID,
	}

	var err error
	tx, err := rsc.db.BeginTX(ctx, nil)
	if err != nil {
		log.Printf("[UpdateUserAccountInDB] rsc.db.BeginTX() got an error: %+v\nMeta: %+v\n", err, meta)
		return err
	}

	defer func() {
		errRollback := rsc.rollbackTX(ctx, tx, err)
		if errRollback != nil {
			log.Printf("[UpdateUserAccountInDB] rsc.rollbackTX() got an error: %+v\nMeta: %+v\n", err, meta)
		}
	}()

	err = rsc.db.UpdateUserAccount(ctx, tx, pgsql.UpdateUserAccountParam(param))
	if err != nil {
		log.Printf("[UpdateUserAccountInDB] rsc.db.UpdateUserAccount() got an error: %+v\nMeta: %+v\n", err, meta)
		return err
	}

	errCommit := rsc.db.Commit(tx)
	if errCommit != nil {
		log.Printf("[UpdateUserAccountInDB] rsc.db.Commit() got an error: %+v\nMeta: %+v\n", errCommit, meta)
	}

	return nil
}

// UpdateUserPasswordInDB will update user's password based on the given parameter.
func (rsc *Resource) UpdateUserPasswordInDB(ctx context.Context, userID int64, password string) error {
	meta := map[string]interface{}{
		"user_id": userID,
	}

	var err error
	tx, err := rsc.db.BeginTX(ctx, nil)
	if err != nil {
		log.Printf("[UpdateUserPasswordInDB] rsc.db.BeginTX() got an error: %+v\nMeta: %+v\n", err, meta)
		return err
	}

	defer func() {
		errRollback := rsc.rollbackTX(ctx, tx, err)
		if errRollback != nil {
			log.Printf("[UpdateUserPasswordInDB] rsc.rollbackTX() got an error: %+v\nMeta: %+v\n", err, meta)
		}
	}()

	err = rsc.db.UpdateUserPassword(ctx, tx, userID, password)
	if err != nil {
		log.Printf("[UpdateUserPasswordInDB] rsc.db.UpdateUserPassword() got an error: %+v\nMeta: %+v\n", err, meta)
		return err
	}

	errCommit := rsc.db.Commit(tx)
	if errCommit != nil {
		log.Printf("[UpdateUserPasswordInDB] rsc.db.Commit() got an error: %+v\nMeta: %+v\n", errCommit, meta)
	}

	return nil
}

// rollbackTX will rollback a transaction if any error occured.
func (rsc *Resource) rollbackTX(ctx context.Context, tx *sql.Tx, err error) error {
	if err == nil {
		return nil
	}

	errRollback := rsc.db.Rollback(tx)
	if errRollback != nil {
		log.Printf("[rollbackTX] rsc.db.Rollback() got an error: %+v\n", err)
		return err
	}

	return nil
}
