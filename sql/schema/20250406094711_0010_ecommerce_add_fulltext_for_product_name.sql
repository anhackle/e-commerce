-- +goose Up
-- +goose StatementBegin
ALTER TABLE `product` ADD FULLTEXT idx_name_desc (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `product` DROP INDEX idx_name_desc;
-- +goose StatementEnd
