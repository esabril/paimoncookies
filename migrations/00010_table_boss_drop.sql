-- +goose Up
CREATE TABLE IF NOT EXISTS boss_drop (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    type text
);

CREATE INDEX idx_character_boss_drop_type ON boss_drop (type);

-- +goose Down
DROP TABLE boss_drop;