-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `orders` (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
    user_id INT NOT NULL,
    payment_method ENUM("COD", "MOMO", "BANK") NOT NULL,
    status ENUM("pending", "paid", "processing", "shipped", "delivered", "cancelled", "failed") DEFAULT "pending",
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    shipping_address TEXT NOT NULL,
    total BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
    
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `orders`;
-- +goose StatementEnd
