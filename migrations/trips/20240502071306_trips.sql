-- +goose Up
-- +goose StatementBegin
CREATE TYPE taxi_type AS ENUM('economy', 'comfort', 'business');
CREATE TABLE IF NOT EXISTS trips(
    id BIGSERIAL PRIMARY KEY,
    from_address VARCHAR(255) NOT NULL,
    to_address VARCHAR(255) NOT NULL,
    rate INT NOT NULL DEFAULT 0,
    taxi taxi_type,
    user_id INT NOT NULL,
    CONSTRAINT fk_user_trips FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS trips DROP CONSTRAINT IF EXISTS fk_taxi_type_trips;
ALTER TABLE IF EXISTS trips DROP CONSTRAINT IF EXISTS fk_user_trips;
DROP TABLE IF EXISTS taxi_types;
DROP TABLE IF EXISTS trips;
-- +goose StatementEnd
