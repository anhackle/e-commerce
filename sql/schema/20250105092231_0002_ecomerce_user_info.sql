-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_info (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
    user_id INT NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    gender BOOLEAN NOT NULL,
    birth_day DATE NOT NULL,
    address TEXT NOT NULL,
    phone VARCHAR(10) NOT NULL,
    description TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_info;
-- +goose StatementEnd
