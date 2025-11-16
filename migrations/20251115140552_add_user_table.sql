-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL,
    team_name VARCHAR(100) NOT NULL REFERENCES team(team_name) ON DELETE CASCADE,
    is_active BOOLEAN NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user";
-- +goose StatementEnd
