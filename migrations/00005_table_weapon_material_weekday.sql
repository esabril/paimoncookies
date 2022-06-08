-- +goose Up
CREATE TABLE IF NOT EXISTS weapon_materials_weekday (
    type text,
    weekday text,
    CONSTRAINT fk_weapon_material_type
        FOREIGN KEY (type)
            REFERENCES weapon_materials_type(type)
);

INSERT INTO weapon_materials_weekday (type, weekday) VALUES
    ('decarabian_tiles', 'monday'),
    ('decarabian_tiles', 'thursday'),
    ('guyun_pillars', 'monday'),
    ('guyun_pillars', 'thursday'),
    ('branches_distant_sea', 'monday'),
    ('branches_distant_sea', 'thursday'),
    ('borean_wolf_teeth', 'tuesday'),
    ('borean_wolf_teeth', 'friday'),
    ('elixirs', 'tuesday'),
    ('elixirs', 'friday'),
    ('narukami_magatamas', 'tuesday'),
    ('narukami_magatamas', 'friday'),
    ('gladiator_shackles', 'wednesday'),
    ('gladiator_shackles', 'saturday'),
    ('aerosiderite', 'wednesday'),
    ('aerosiderite', 'saturday'),
    ('oni_masks', 'wednesday'),
    ('oni_masks', 'saturday')
;


-- +goose Down
DROP TABLE weapon_materials_weekday;

