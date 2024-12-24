package objecttypesmigrations

const createUserTypeENUM = `
	CREATE TYPE public.user_type AS enum (
		'client', 
		'restaurant_admin', 
		'restaurant_worker', 
		'admin'
	);
	
	ALTER TYPE public.user_type OWNER TO discount_club;
`

const dropUserTypeENUM = `
	DROP TYPE IF EXISTS public.user_type CASCADE;
`
