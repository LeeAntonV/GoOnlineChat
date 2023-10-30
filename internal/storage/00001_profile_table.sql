-- +goose Up
-- +goose StatementBegin
CREATE TABLE profile
(
    ID          SERIAL       PRIMARY KEY,
    Name        TEXT         NOT NULL,
    PASSWORD    TEXT         NOT NULL,
    FriendsWith TEXT         NOT NULL ,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profile;
-- +goose StatementEnd
