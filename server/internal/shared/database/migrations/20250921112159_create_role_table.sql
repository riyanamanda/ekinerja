-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(20) NOT NULL UNIQUE
);
CREATE UNIQUE INDEX idx_role_nama ON role(nama);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role;
-- +goose StatementEnd
