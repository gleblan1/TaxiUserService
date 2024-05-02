-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS trips(
    id BIGSERIAL PRIMARY KEY,
    from_address VARCHAR(255) NOT NULL,
    to_address VARCHAR(255) NOT NULL,
    rate INT NOT NULL DEFAULT 0,
    taxi_type INT,
    user_id INT NOT NULL,
    CONSTRAINT fk_user_trips FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS taxi_types(
     id SERIAL PRIMARY KEY,
     type ENUM('economy', 'comfort', 'business')
);

INSERT INTO taxi_types (type) VALUES ('economy'), ('comfort'), ('business');

ALTER TABLE taxi_types ALTER COLUMN type SET DEFAULT 'economy';

ALTER TABLE IF EXISTS trips ADD CONSTRAINT fk_taxi_type_trips FOREIGN KEY (taxi_type) REFERENCES taxi_types(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS trips DROP CONSTRAINT IF EXISTS fk_taxi_type_trips;
ALTER TABLE IF EXISTS trips DROP CONSTRAINT IF EXISTS fk_user_trips;
DROP TABLE IF EXISTS taxi_types;
DROP TABLE IF EXISTS trips;
-- +goose StatementEnd
