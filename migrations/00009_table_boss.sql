-- +goose Up
CREATE TABLE IF NOT EXISTS weekly_boss (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    location text,
    domain text,
    talent_materials text [],
    gems text []
);

CREATE INDEX idx_weekly_boss_location ON weekly_boss (location);

CREATE TABLE IF NOT EXISTS world_boss (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    location text,
    ascension_material text,
    gems text []    
);

CREATE INDEX idx_world_boss_location ON world_boss (location);

-- +goose Down
DROP TABLE weekly_boss;
DROP TABLE world_boss;
