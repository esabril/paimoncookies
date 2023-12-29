-- +goose Up
CREATE TABLE IF NOT EXISTS talent_book_type (
    id SERIAL PRIMARY KEY,
    title text NOT NULL,
    type text UNIQUE NOT NULL,
    location text
);

CREATE INDEX idx_talent_book_type ON talent_book_type (
    type
);

CREATE INDEX idx_talent_book_location ON talent_book_type (
    location
);

-- +goose Down
DROP TABLE talent_book_type;
