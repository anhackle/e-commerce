-- +goose Up
-- +goose StatementBegin
ALTER TABLE product ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product DROP COLUMN deleted_at;
-- +goose StatementEnd
