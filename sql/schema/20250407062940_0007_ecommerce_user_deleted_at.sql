-- +goose Up
-- +goose StatementBegin
ALTER TABLE user ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user DROP COLUMN deleted_at;
-- +goose StatementEnd
