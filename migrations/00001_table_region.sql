-- +goose Up
CREATE TABLE IF NOT EXISTS region (
    id SERIAL PRIMARY KEY,
    name text UNIQUE NOT NULL,
    title text
);

CREATE INDEX idx_name ON region (
    name
);

INSERT INTO region (name, title) VALUES
    ('mondstadt', 'Мондштадт'),
    ('liyue', 'Ли Юэ'),
    ('inazuma', 'Инадзума')
;

-- +goose Down
DROP TABLE region;
