package objecttypesmigrations

const createObjectTypeENUM = `
	CREATE TYPE user_type AS ENUM (
		'client', 
		'restaurant_admin', 
		'restaurant_worker', 
		'admin'
	);
`