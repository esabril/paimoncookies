-- +goose Up
CREATE TABLE IF NOT EXISTS ascension_material (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    type text,
    location text []
);

CREATE INDEX idx_ca_type ON ascension_material (type);

-- +goose Down
DROP TABLE ascension_material;
