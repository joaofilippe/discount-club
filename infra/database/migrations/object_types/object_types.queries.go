package objecttypesmigrations

const createUserTypeENUM = `
	CREATE TYPE user_type AS ENUM (
		'client', 
		'restaurant_admin', 
		'restaurant_worker', 
		'admin'
	);
`
