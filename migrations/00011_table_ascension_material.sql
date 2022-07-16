-- +goose Up
CREATE TABLE IF NOT EXISTS ascension_material (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    type text,
    location text []
);

CREATE INDEX idx_ca_type ON ascension_material (
    type
);

INSERT INTO ascension_material (name, title, type, location) VALUES
    ('valberry', 'Валяшка', 'local_speciality', '{"mondstadt"}'),
    ('windwheel_aster', 'Ветряная астра', 'local_speciality', '{"mondstadt"}'),
    ('wolfhook', 'Волчий крюк', 'local_speciality', '{"mondstadt"}'),
    ('philanemo_mushroom', 'Гриб филанемо', 'local_speciality', '{"mondstadt"}'),
    ('calla_lily', 'Лилия калла', 'local_speciality', '{"mondstadt"}'),
    ('dandelion_seed', 'Семена одуванчика', 'local_speciality', '{"mondstadt"}'),
    ('cecilia', 'Сесилия', 'local_speciality', '{"mondstadt"}'),
    ('small_lamp_grass', 'Трава-светяшка', 'local_speciality', '{"mondstadt"}'),
    ('glaze_lily', 'Глазурная лилия', 'local_speciality', '{"liyue"}'),
    ('jueyun_chili', 'Заоблачный перчик', 'local_speciality', '{"liyue"}'),
    ('starconch', 'Звёздная ракушка', 'local_speciality', '{"liyue"}'),
    ('cor_lapis', 'Кор ляпис', 'local_speciality', '{"liyue"}'),
    ('noctilucous_jade', 'Полуночный нефрит', 'local_speciality', '{"liyue"}'),
    ('violetgrass', 'Стеклянные колокольчики', 'local_speciality', '{"liyue"}'),
    ('qingxin', 'Цветок цинсинь', 'local_speciality', '{"liyue"}'),
    ('silk_flower', 'Шелковица', 'local_speciality', '{"liyue"}'),
    ('sango_pearl', 'Жемчужина Санго', 'local_speciality', '{"inazuma"}'),
    ('crystal_marrow', 'Кристальный костный мозг', 'local_speciality', '{"inazuma"}'),
    ('dendrobium', 'Кровоцвет', 'local_speciality', '{"inazuma"}'),
    ('sea_ganoderma', 'Морской гриб', 'local_speciality', '{"inazuma"}'),
    ('onikabuto', 'Оникабуто', 'local_speciality', '{"inazuma"}'),
    ('amakumo_fruit', 'Плод облачной травы', 'local_speciality', '{"inazuma"}'),
    ('fluorescent_fungus', 'Светящийся гриб', 'local_speciality', '{"inazuma"}'),
    ('naku_weed', 'Трава наку', 'local_speciality', '{"inazuma"}'),
    ('sakura_bloom', 'Цвет сакуры', 'local_speciality', '{"inazuma"}'),
    ('slime', 'Слаймы', 'common', '{"mondstadt", "liyue", "inazuma"}'),
    ('hilichurl_masks', 'Маски хиличурлов', 'common', '{"mondstadt", "liyue", "inazuma"}'),
    ('hilichurl_arrowheads', 'Наконечники стрел хиличурлов', 'common', '{"mondstadt", "liyue", "inazuma"}'),
    ('samachurl_scrolls', 'Свитки шамачурлов', 'common', '{"mondstadt", "liyue", "inazuma"}'),
    ('treasure_hoarder_insignias', 'Печати похитителей сокровищ', 'common', '{"mondstadt", "liyue", "inazuma"}'),
    ('fatui_insignia', 'Шевроны Фатуи', 'common', '{"mondstadt", "liyue", "inazuma"}'),
    ('whopperflower_nectar', 'Нектары попрыгуний', 'common', '{"mondstadt", "liyue", "inazuma"}'),
    ('nobushi_handguards', 'Гарды Нобуси', 'common', '{"inazuma"}'),
    ('spectral_cores', 'Призраки', 'common', '{"inazuma"}'),
    ('fungal_spore_powder', 'Порошок плесенников', 'common', '{"liyue"}')
;

-- +goose Down
DROP TABLE ascension_material;
