-- +goose Up
CREATE TABLE IF NOT EXISTS element(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text NOT NULL
);

-- +goose Down
DROP TABLE element;
