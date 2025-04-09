-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `order_item` (
    id CHAR(36) PRIMARY KEY,
    order_id CHAR(36) NOT NULL,
    name VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    description TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    price BIGINT NOT NULL,
    quantity INT NOT NULL,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (order_id) REFERENCES `orders`(id)
    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `order_item`;
-- +goose StatementEnd
