-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE posts
-- +goose StatementEnd