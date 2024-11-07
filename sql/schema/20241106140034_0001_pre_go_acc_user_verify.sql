-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_acc_user_verify_9999 (
    verify_id INT AUTO_INCREMENT PRIMARY KEY,
    verify_otp VARCHAR(6) NOT NULL,
    verify_key VARCHAR(255) NOT NULL,
    verify_key_hash VARCHAR(255) NOT NULL,
    verify_type INT DEFAULT 1,
    is_verified INT DEFAULT 0,
    is_deleted INT DEFAULT 0,
    verify_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    verify_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_verify_opt (verify_otp),
    UNIQUE KEY unique_verify_key (verify_key)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_acc_user_verify_9999;
-- +goose StatementEnd
