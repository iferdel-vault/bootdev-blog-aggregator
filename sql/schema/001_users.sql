-- +goose Up
CREATE TABLE users (
	id UUID PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	name VARCHAR(50) NOT NULL UNIQUE
);
-- +goose Down
DROP TABLE users;
