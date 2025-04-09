-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `orders` (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    payment_method ENUM("COD", "MOMO", "BANK") NOT NULL,
    status ENUM("pending", "paid", "processing", "shipped", "delivered", "cancelled", "failed") DEFAULT "pending",
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    shipping_address TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    total BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `orders`;
-- +goose StatementEnd
