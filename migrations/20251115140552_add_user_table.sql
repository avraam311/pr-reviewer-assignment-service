-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS user (
    user_id TEXT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    team_name VARCHAR(100) NOT NULL REFERENCES team(team_name) ON DELETE CASCADE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE member;
-- +goose StatementEnd
