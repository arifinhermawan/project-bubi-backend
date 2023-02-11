package pgsql

import (
	// golang package
	"context"
	"database/sql"
	"log"
	"time"
)

// GetUserAccountByEmail will fetch user's information based of account's email.
func (repo *DBRepository) GetUserAccountByEmail(ctx context.Context, email string) (Account, error) {
	infra := repo.infra
	timeout := time.Duration(infra.GetConfig().Database.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	namedParam := map[string]interface{}{
		"email": email,
	}

	namedQuery, args, err := funcSQLXNamed(queryGetUserAccountByEmail, namedParam)
	if err != nil {
		log.Printf("[GetUserAccountByEmail] funcSQLXNamed got an error: %+v\nMeta:%+v\n", err, namedParam)
		return Account{}, err
	}

	var result Account
	err = repo.db.GetContext(ctxQuery, &result, repo.db.Rebind(namedQuery), args...)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[GetUserAccountByEmail] repo.db.GetContext() got an error: %+v\nMeta:%+v\n", err, namedParam)
		return Account{}, err
	}

	return result, nil
}

// InsertUserAccount will create a new entry in table user_account in database.
func (repo *DBRepository) InsertUserAccount(ctx context.Context, tx *sql.Tx, email, password string) error {
	timeout := time.Duration(repo.infra.GetConfig().Database.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	namedParam := map[string]interface{}{
		"email":      email,
		"password":   password,
		"created_at": repo.infra.GetTimeGMT7(),
	}

	meta := map[string]interface{}{
		"email": email,
	}

	namedQuery, args, err := funcSQLXNamed(queryInsertUserAccount, namedParam)
	if err != nil {
		log.Printf("[InsertUserAccount] funcSQLXNamed got an error: %+v\nMeta:%+v\n", err, meta)
		return err
	}

	_, errQuery := tx.ExecContext(ctxQuery, repo.db.Rebind(namedQuery), args...)
	if errQuery != nil {
		log.Printf("[InsertUserAccount] tx.ExecContext got an error: %+v\nMeta:%+v\n", errQuery, meta)
		return errQuery
	}

	return nil
}

// UpdateUserAccount will update user's account information.
func (repo *DBRepository) UpdateUserAccount(ctx context.Context, tx *sql.Tx, param UpdateUserAccountParam) error {
	timeout := time.Duration(repo.infra.GetConfig().Database.DefaultTimeout) * time.Second
	ctxQuery, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	namedParam := map[string]interface{}{
		"first_name":    param.FirstName,
		"id":            param.UserID,
		"last_name":     param.LastName,
		"record_period": param.RecordPeriod,
		"updated_at":    repo.infra.GetTimeGMT7(),
	}

	namedQuery, args, err := funcSQLXNamed(queryUpdateUserAccount, namedParam)
	if err != nil {
		log.Printf("[UpdateUserAccount] funcSQLXNamed got an error: %+v\nMeta:%+v\n", err, namedParam)
		return err
	}

	_, err = tx.ExecContext(ctxQuery, repo.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Printf("[UpdateUserAccount] tx.ExecContext() got an error: %+v\nMeta:%+v\n", err, namedParam)
		return err
	}

	return nil
}
