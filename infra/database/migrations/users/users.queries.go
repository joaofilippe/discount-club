package usersmigrations

const createTableQuery = `
	CREATE TABLE IF NOT EXISTS users(
		id UUID PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		user_type user_type NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP,
		active BOOLEAN NOT NULL
	)
`

const dropTableQuery = `
	DROP TABLE IF EXISTS users;
`