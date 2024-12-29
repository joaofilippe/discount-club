CREATE TYPE user_type AS enum ('client', 'restaurant_admin', 'restaurant_worker', 'admin');

CREATE TABLE
    IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        email VARCHAR(255) NOT NULL,
        PASSWORD VARCHAR(255) NOT NULL,
        user_type user_type NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP,
        active BOOLEAN NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS discounts (
        id UUID NOT NULL PRIMARY KEY,
        restaurant_id UUID NOT NULL,
        user_id UUID NOT NULL,
        code VARCHAR(255) NOT NULL,
        percentage INTEGER NOT NULL,
        start_date TIMESTAMP NOT NULL,
        end_date TIMESTAMP NOT NULL,
        times INTEGER NOT NULL,
        description TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP,
        active BOOLEAN NOT NULL
    );

ALTER TABLE users OWNER TO discount_club;

ALTER TABLE discounts OWNER TO discount_club;

ALTER TYPE user_type OWNER TO discount_club;