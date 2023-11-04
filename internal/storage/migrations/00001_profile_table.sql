-- +goose Up
-- +goose StatementBegin
CREATE TABLE profile
(
    ID          SERIAL       PRIMARY KEY ,
    EMAIL       VARCHAR(255) UNIQUE ,
    HASH    VARCHAR(255) NOT NULL,
    created_at  DATE    DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profile;
-- +goose StatementEnd
