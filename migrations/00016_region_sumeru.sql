-- +goose Up
INSERT INTO world_boss (name, title, location, ascension_material, gems)
VALUES ('electro_regisvine', 'Электро Папоротник', 'sumeru', 'thunderclap_fruitcore', '{"vajrada_amethyst"}');

INSERT INTO boss_drop (name, title, type)
VALUES ('thunderclap_fruitcore', 'Плод грома', 'ascension');

INSERT INTO ascension_material (name, title, type, location)
VALUES ('kalpalata_lotus', 'Лотос кальпалата', 'local_speciality', '{sumeru}'),
       ('red_satin', 'Красный шелк', 'common', '{sumeru}');

-- +goose Down
DELETE FROM world_boss WHERE name = 'electro_regisvine';
DELETE FROM boss_drop WHERE name = 'thunderclap_fruitcore';
DELETE FROM ascension_material WHERE name = 'kalpalata_lotus';
