-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bidang (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(150) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_bidang_nama_unique ON bidang(nama);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bidang;
-- +goose StatementEnd
