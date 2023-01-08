package mysql

func Query() (query []string) {
	account := `
	CREATE TABLE IF NOT EXISTS account (
		id integer PRIMARY KEY AUTO_INCREMENT,
		name varchar(19) NOT NULL,
		email varchar(29) NOT NULL UNIQUE,
		password varchar(60) NOT NULL,
		created_at datetime DEFAULT CURRENT_TIMESTAMP
	);		
	`

	query = append(query, account)
	return
}
