-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS jabatan (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);
CREATE INDEX IF NOT EXISTS idx_jabatan_nama_unique ON jabatan(nama);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS jabatan;
-- +goose StatementEnd
