-- +goose Up
CREATE TABLE IF NOT EXISTS talent_books_type (
    id SERIAL PRIMARY KEY,
    title text NOT NULL,
    type text UNIQUE NOT NULL,
    location text
);

CREATE INDEX idx_talent_books_type ON talent_books_type (
    type
);

CREATE INDEX idx_talent_books_location ON talent_books_type (
    location
);

INSERT INTO talent_books_type (title, type, location) VALUES
    ('О Свободе', 'freedom', 'mondstadt'),
    ('О Борьбе', 'resistance', 'mondstadt'),
    ('О Поэзии', 'ballad', 'mondstadt'),
    ('О Процветании', 'prosperity', 'liyue'),
    ('Об Усердии', 'diligence', 'liyue'),
    ('О Золоте', 'gold', 'liyue'),
    ('О Бренности', 'transience', 'inazuma'),
    ('Об изяществе', 'elegance', 'inazuma'),
    ('О свете', 'light', 'inazuma')
;

-- +goose Down
DROP TABLE talent_books_type;
