-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_acc_user_info_9999 (
    user_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_account VARCHAR(255) NOT NULL,
    user_nickname VARCHAR(255),
    user_avatar VARCHAR(255),
    user_state TINYINT UNSIGNED NOT NULL,
    user_mobile VARCHAR(20),

    user_gender TINYINT UNSIGNED,
    user_birthday DATE,
    user_email VARCHAR(255),
    
    user_is_authentication TINYINT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE KEY unique_user_account (user_account),
    INDEX idex_user_mobile (user_mobile),
    INDEX idx_user_email (user_email),
    INDEX idx_user_state (user_state),
    INDEX idx_user_is_authentication (user_is_authentication)

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_info_9999`;
-- +goose StatementEnd
