-- +goose Up
-- +goose StatementBegin
ALTER TABLE `cart` ADD UNIQUE KEY unique_user_product (user_id, product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `cart` DROP INDEX unique_user_product;
-- +goose StatementEnd
