package pgsql

const (
	queryGetUserAccountByEmail = `
		SELECT 
			email, 
			record_period_start, 
			first_name, 
			last_name, 
			id,
			password
		FROM
			user_account
		WHERE
			email = :email
	`

	queryInsertUserAccount = `
		INSERT INTO 
			user_account(email,"password",created_at)
		VALUES (
			:email,
			:password,
			:created_at
		)
	`

	queryUpdateUserAccount = `
		UPDATE
			user_account
		SET 
			first_name = :first_name,
			last_name = :last_name,
			record_period_start = :record_period,
			updated_at = :updated_at
		WHERE
			id = :id
	`
)
