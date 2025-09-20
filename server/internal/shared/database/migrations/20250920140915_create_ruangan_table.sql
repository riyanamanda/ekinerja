-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ruangan (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_ruangan_nama ON ruangan (nama);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ruangan;
-- +goose StatementEnd
