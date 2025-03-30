-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `order_item` (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
    order_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price BIGINT NOT NULL,
    quantity INT NOT NULL,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (order_id) REFERENCES `orders`(id)
)CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `order_item`;
-- +goose StatementEnd
