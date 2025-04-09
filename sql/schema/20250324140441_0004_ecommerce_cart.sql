-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cart (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    product_id CHAR(36) NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (product_id) REFERENCES product(id),
    UNIQUE INDEX unique_user_product (user_id, product_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart;
-- +goose StatementEnd
