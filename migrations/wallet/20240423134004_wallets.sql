-- +goose Up
-- +goose StatementBegin
CREATE TABLE wallets(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    is_family BOOLEAN NOT NULL DEFAULT FALSE,
    balance BIGINT DEFAULT 0,
    deleted_at timestamptz,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

ALTER TABLE wallets ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade;

CREATE TABLE transactions(
    id BIGSERIAL PRIMARY KEY,
    from_wallet BIGINT NOT NULL,
    to_wallet BIGINT NOT NULL,
    amount FLOAT NOT NULL,
    status VARCHAR(255) NOT NULL
);

ALTER TABLE transactions ADD CONSTRAINT fk_from_wallet FOREIGN KEY (from_wallet) REFERENCES wallets (id) on delete cascade;

ALTER TABLE transactions ADD CONSTRAINT fk_to_wallet  FOREIGN KEY (to_wallet) REFERENCES wallets (id) on delete cascade;

CREATE TABLE IF NOT EXISTS family_wallets(
    wallet_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    is_owner BOOLEAN NOT NULL DEFAULT FALSE
);

ALTER TABLE family_wallets ADD CONSTRAINT fk_family_user_id FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE family_wallets ADD CONSTRAINT fk_family_wallet_id FOREIGN KEY (wallet_id) REFERENCES wallets (id);

ALTER TABLE users ADD CONSTRAINT fk_wallet_id FOREIGN KEY (current_wallet_id) REFERENCES wallets (id) on delete cascade;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions DROP CONSTRAINT fk_from_wallet;
ALTER TABLE transactions DROP CONSTRAINT fk_to_wallet;
ALTER TABLE family_wallets DROP CONSTRAINT fk_family_user_id;
ALTER TABLE family_wallets DROP CONSTRAINT fk_family_wallet_id;
ALTER TABLE users DROP CONSTRAINT fk_wallet_id;
DROP TABLE transactions;
DROP TABLE wallets;
-- +goose StatementEnd
