-- +goose Up
CREATE TABLE IF NOT EXISTS character(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text NOT NULL,
    region text NOT NULL,
    rarity int8 NOT NULL,
    element text NOT NULL,
    talent_book_type text,
    talent_boss_drop text,
    ascension_boss_drop text,
    ascension_gem text,
    ascension_local_speciality text,
    common_ascension_material text    
);

CREATE INDEX idx_character_name ON character (title);
CREATE INDEX idx_character_element ON character (element);
CREATE INDEX idx_character_talent_book_type ON character (talent_book_type);

-- +goose Down
DROP TABLE character;
