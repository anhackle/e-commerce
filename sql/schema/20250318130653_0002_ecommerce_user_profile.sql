-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_profile (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    first_name VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    last_name VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    phone_number VARCHAR(15),
    address TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profile;
-- +goose StatementEnd
