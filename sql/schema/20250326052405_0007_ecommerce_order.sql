-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `order` (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
    user_id INT NOT NULL,
    payment_method ENUM("COD", "MOMO", "BANK"),
    status ENUM("create", "confirm", "pay", "ship", "finish") DEFAULT "create",
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    shipping_address TEXT,
    total INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
)CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `order`;
-- +goose StatementEnd
