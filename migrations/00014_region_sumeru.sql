-- +goose Up
INSERT INTO region (name, title) VALUES ('sumeru', 'Сумеру');
INSERT INTO gem (name, title) VALUES ('nagadus_emerald', 'Изумруд Нагадус');
INSERT INTO world_boss (name, title, location, ascension_material, gems) VALUES (
    'jadeplume_terrorshroom', 'Пернатый плесенник', 'sumeru', 'majestic_hooked_beak', '{"nagadus_emerald"}'
);
INSERT INTO boss_drop (name, title, type) VALUES (
    'majestic_hooked_beak', 'Клюв короля плесенников', 'ascension'
);
INSERT INTO ascension_material (name, title, type, location) VALUES
    ('nilotpala_lotus', 'Лотос нилотпала', 'local_speciality', '{sumeru}'),
    ('rukkhashava_mushrooms', 'Грибы руккхашава', 'local_speciality', '{sumeru}')
;

UPDATE ascension_material SET location = array_append(location, 'sumeru') WHERE name = 'fungal_spore_powder';

INSERT INTO talent_books_type (title, type, location) VALUES
    ('О Наставлениях', 'admonition', 'sumeru'),
    ('Об Остроумии', 'ingenuity', 'sumeru'),
    ('О Честности', 'praxis', 'sumeru')
;
INSERT INTO talent_books_weekday (type, weekday) VALUES
    ('admonition', 'monday'),
    ('admonition', 'thursday'),
    ('ingenuity', 'tuesday'),
    ('ingenuity', 'friday'),
    ('praxis', 'wednesday'),
    ('praxis', 'saturday')
;

INSERT INTO weapon_materials_type (title, alias, type, location) VALUES
    ('Амулет лесной росы', '', 'talisman_forest_dew', 'sumeru'),
    ('Цветущий оазис', 'тарелки', 'oasis_gardens_reminiscence', 'sumeru'),
    ('Пылающей мощь', '', 'echo_scorching_might', 'sumeru')
;
INSERT INTO weapon_materials_weekday (type, weekday) VALUES
    ('talisman_forest_dew', 'monday'),
    ('talisman_forest_dew', 'thursday'),
    ('oasis_gardens_reminiscence', 'tuesday'),
    ('oasis_gardens_reminiscence', 'friday'),
    ('echo_scorching_might', 'wednesday'),
    ('echo_scorching_might', 'saturday')
;

-- +goose Down
DELETE FROM region WHERE name = 'sumeru';
DELETE FROM gem WHERE name = 'nagadus_emerald';
DELETE FROM world_boss WHERE name = 'jadeplume_terrorshroom';
DELETE FROM boss_drop WHERE name = 'majestic_hooked_beak';
DELETE FROM ascension_material WHERE name IN ('nilotpala_lotus', 'rukkhashava_mushrooms');
UPDATE ascension_material SET location = array_remove(location, 'sumeru') WHERE name = 'fungal_spore_powder';
DELETE FROM talent_books_type WHERE type IN ('admonition', 'ingenuity', 'praxis');
DELETE FROM weapon_materials_type WHERE type IN('talisman_forest_dew', 'oasis_gardens_reminiscence', 'echo_scorching_might');