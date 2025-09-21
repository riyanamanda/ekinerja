-- +goose Up
-- +goose StatementBegin
INSERT INTO role (nama) VALUES
('admin'),
('user');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM role WHERE nama IN ('admin', 'user');
-- +goose StatementEnd
