-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS team (
    id SERIAL PRIMARY KEY,
    team_name VARCHAR(100) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS team;
-- +goose StatementEnd
