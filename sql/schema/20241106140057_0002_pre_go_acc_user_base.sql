-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_acc_user_base_9999 (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    user_account VARCHAR(255) NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    user_salt VARCHAR(255) NOT NULL,

    user_login_time TIMESTAMP NULL DEFAULT NULL,
    user_logout_time TIMESTAMP NULL DEFAULT NULL,
    user_login_ip VARCHAR(45) NULL,

    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY unique_user_account (user_account)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
-- +goose StatementEnd
