-- +goose Up
INSERT INTO character (
    name, title, region, rarity, element, talent_book_type, talent_boss_drop,
    ascension_boss_drop, ascension_gem, ascension_local_speciality, common_ascension_material
) VALUES
    ('dori', 'Дори', 'sumeru', 4, 'electro', 'ingenuity', 'bloodjade_branch',
     'thunderclap_fruitcore', 'vajrada_amethyst', 'kalpalata_lotus', 'red_satin');

-- +goose Down
DELETE FROM character WHERE name = 'dori';
