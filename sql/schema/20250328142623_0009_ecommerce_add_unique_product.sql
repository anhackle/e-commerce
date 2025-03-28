-- +goose Up
-- +goose StatementBegin
ALTER TABLE `product` ADD UNIQUE KEY unique_product (name,price);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `product` DROP INDEX unique_product;
-- +goose StatementEnd
