package discountmigrations

const createTableQuery = `
	CREATE TABLE IF NOT EXISTS discount (
		id UUID PRIMARY KEY,
		restaurant_id UUID NOT NULL,
		user_id UUID NOT NULL,
		code VARCHAR(255) NOT NULL,
		percentage INT NOT NULL,
		start_date TIMESTAMP NOT NULL,
		end_date TIMESTAMP NOT NULL,
		times INT NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP,
		active BOOLEAN NOT NULL
	);
`

const dropTableQuery = `
	DROP TABLE IF EXISTS discount;
`
