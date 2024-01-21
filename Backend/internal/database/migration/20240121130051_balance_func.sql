-- +goose Up
-- +goose StatementBegin
ALTER TABLE "User"
DROP COLUMN balance;

CREATE TABLE "Balance" (
    id TEXT PRIMARY KEY,
    user_id TEXT,
    balance FLOAT,
    balance_name TEXT,
    FOREIGN KEY (user_id) REFERENCES "User"(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "User"
ADD COLUMN balance FLOAT; 
DROP TABLE IF EXISTS Balance;
-- +goose StatementEnd
