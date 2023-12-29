-- +goose Up
CREATE TABLE IF NOT EXISTS gem (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text
);

-- +goose Down
DROP TABLE gem;
