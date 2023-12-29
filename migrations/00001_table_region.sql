-- +goose Up
CREATE TABLE IF NOT EXISTS region (
    id SERIAL PRIMARY KEY,
    name text UNIQUE NOT NULL,
    title text
);

CREATE INDEX idx_name ON region (
    name
);

-- +goose Down
DROP TABLE region;
