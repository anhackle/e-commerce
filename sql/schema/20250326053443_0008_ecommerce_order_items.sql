-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `order_item` (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    present_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (order_id) REFERENCES `order`(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `order_item`;
-- +goose StatementEnd
