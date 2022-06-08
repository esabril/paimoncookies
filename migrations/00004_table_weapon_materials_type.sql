-- +goose Up
CREATE TABLE IF NOT EXISTS weapon_materials_type (
    id SERIAL PRIMARY KEY,
    title text NOT NULL,
    alias text,
    type text UNIQUE NOT NULL,
    location text NOT NULL
);

CREATE INDEX idx_weapon_materials_type ON weapon_materials_type (
    type
);

CREATE INDEX idx_weapon_materials_location ON weapon_materials_type (
    location
);

INSERT INTO weapon_materials_type (title, alias, type, location) VALUES
    ('Плитки Декарабиана', '', 'decarabian_tiles', 'mondstadt'),
    ('Клыки арктического волка', '', 'borean_wolf_teeth', 'mondstadt'),
    ('Кандалы Гладиатора', '', 'gladiator_shackles', 'mondstadt'),
    ('Столбы Гуюнь', '', 'guyun_pillars', 'liyue'),
    ('Пилюли Заоблачного моря', 'жемчужины', 'elixirs', 'liyue'),
    ('Черный сидерит', 'трубки', 'aerosiderite', 'liyue'),
    ('Ветви внешних морей', '', 'branches_distant_sea', 'inazuma'),
    ('Бусины Наруками', 'запятые', 'narukami_magatamas', 'inazuma'),
    ('Маски Они', '', 'oni_masks', 'inazuma')
;
-- +goose Down
DROP TABLE weapon_materials_type;

