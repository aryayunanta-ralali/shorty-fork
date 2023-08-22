-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS short_urls
(
    id          BIGINT(20) AUTO_INCREMENT PRIMARY KEY,
    user_id     VARCHAR(255) NULL,
    url         VARCHAR(255) NOT NULL,
    short_code  VARCHAR(255) NOT NULL UNIQUE,
    visit_count INT UNSIGNED NOT NULL DEFAULT 0,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),

    KEY user_id (user_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS short_urls;
-- +goose StatementEnd
