package mysql

var CreateAccountState = `
	insert into
		account(
			name,
			email,
			password
		)
		values(?, ?, ?)
`

var FindAccountState = `
		select
			id,
			name,
			email,
			password,
			created_at
		from
			account
		where
			id = ?
`
