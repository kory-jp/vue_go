package mysql

var GetAccountByEmail = `
	select
		id,
		name,
		email,
		password,
		created_at
	from
		account
	where
		email = ?
`
