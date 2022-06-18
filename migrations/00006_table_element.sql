-- +goose Up
CREATE TABLE IF NOT EXISTS element(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text NOT NULL
);

INSERT INTO element (name, title) VALUES
    ('cryo', 'Крио'),
    ('pyro', 'Пиро'),
    ('electro', 'Электро'),
    ('anemo', 'Анемо'),
    ('hydro', 'Гидро'),
    ('geo', 'Гео'),
    ('dendro', 'Дендро')
;

-- +goose Down
DROP TABLE element;
