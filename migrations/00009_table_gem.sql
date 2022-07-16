-- +goose Up
CREATE TABLE IF NOT EXISTS gem (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text
);

INSERT INTO gem (name, title) VALUES
    ('agnidus_agate', 'Агат Агнидус'),
    ('varunada_lazurite', 'Лазурит Варунада'),
    ('vajrada_amethyst', 'Аметист Ваджрада'),
    ('vayuda_turquoise', 'Бирюза Вайюда'),
    ('shivada_jade', 'Нефрит Шивада'),
    ('prithiva_topaz', 'Топаз Притхива')
;

-- +goose Down
DROP TABLE gem;
