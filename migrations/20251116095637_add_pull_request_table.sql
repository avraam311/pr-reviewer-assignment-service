-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pull_request (
    id SERIAL PRIMARY KEY,
    pull_request_id TEXT NOT NULL,
    pull_request_name VARCHAR(255) NOT NULL,
    author_id TEXT NOT NULL,
    status VARCHAR(10) NOT NULL,
    assigned_reviewers TEXT[] NOT NULL,
    merged_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pull_request;
-- +goose StatementEnd
