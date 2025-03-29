-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_profile (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
    user_id INT NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(15),
    address VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES user(id)
)CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profile;
-- +goose StatementEnd
