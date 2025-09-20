-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS atasan (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) UNIQUE NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_atasan_nama_unique ON atasan (nama);
CREATE INDEX IF NOT EXISTS idx_atasan_is_active ON atasan (is_active);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS atasan;
-- +goose StatementEnd
