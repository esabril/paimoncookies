-- +goose Up
CREATE TABLE IF NOT EXISTS weekly_boss (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    location text,
    domain text,
    talent_materials text [],
    gems text []
);

CREATE INDEX idx_weekly_boss_location ON weekly_boss (
    location
);

INSERT INTO weekly_boss (name, title, location, domain, talent_materials, gems) VALUES
    ('andrius', 'Андриус', 'mondstadt', 'Испытание Волка Севера', '{"tail_of_boreas", "ring_of_boreas", "spirit_locker_of_boreas"}', '{"shivada_jade", "prithiva_topaz", "agnidus_agate"}'),
    ('dvalin', 'Двалин', 'mondstadt', 'Логово Ужаса Бури', '{"dvalins_plume", "dvalins_claw", "dvalins_sigh"}', '{"vayuda_turquoise", "vajrada_amethyst", "varunada_lazurite"}'),
    ('childe', 'Чайльд', 'liyue', 'Золотая палата', '{"tusk_of_monoceros_caeli", "shard_of_foul_legacy", "shadow_of_the_warrior"}', '{"varunada_lazurite", "vajrada_amethyst", "shivada_jade"}'),
    ('azhdaha', 'Аждаха', 'liyue', 'Под древом подавления', '{"dragon_lords_crown", "bloodjade_branch", "gilded_scale"}', '{"agnidus_agate", "varunada_lazurite", "vajrada_amethyst", "shivada_jade", "prithiva_topaz"}'),
    ('la_signora', 'Синьора', 'inazuma', 'Остров Наруками: Тэнсюкаку', '{"molten_moment", "hellfire_butterfly", "ashen_heart"}', '{"shivada_jade", "agnidus_agate"}'),
    ('magatsu_mitake_narukami_no_mikoto', 'Магацу митакэ Наруками но микото', 'inazuma', 'Конец Царства онейроса', '{"mudra_of_the_malefic_general", "tears_of_the_calamitous_god", "the_meaning_of_aeons"}', '{"vajrada_amethyst"}')
;

CREATE TABLE IF NOT EXISTS world_boss (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    location text,
    ascension_material text,
    gems text []
);

CREATE INDEX idx_world_boss_location ON world_boss (
    location
);

INSERT INTO world_boss (name, title, location, ascension_material, gems) VALUES
    ('anemo_hypostasis', 'Анемо гипостазис', 'mondstadt', 'hurricane_seed', '{"vayuda_turquoise"}'),
    ('electro_hypostasis', 'Электро гипостазис', 'mondstadt', 'lightning_prism', '{"vajrada_amethyst"}'),
    ('geo_hypostasis', 'Гео гипостазис', 'liyue', 'basalt_pillar', '{"prithiva_topaz"}'),
    ('cryo_regisvine', 'Крио папоротник', 'mondstadt', 'hoarfrost_core', '{"shivada_jade"}'),
    ('pyro_regisvine', 'Пиро орхидея', 'liyue', 'everflame_seed', '{"agnidus_agate"}'),
    ('rhodeia_of_loch', 'Океанид', 'liyue', 'cleansing_heart', '{"varunada_lazurite"}'),
    ('primo_geovishap', 'Древний геовишап', 'liyue', 'juvenile_jade', '{"prithiva_topaz", "vajrada_amethyst", "shivada_jade", "agnidus_agate", "varunada_lazurite"}'),
    ('cryo_hypostasis', 'Крио гипостазис', 'mondstadt', 'crystalline_bloom', '{"shivada_jade"}'),
    ('maguu_kenki', 'Магу кэнки', 'inazuma', 'marionette_core', '{"shivada_jade", "vayuda_turquoise"}'),
    ('perpetual_mechanical_array', 'Бесконечный механический массив', 'inazuma', 'perpetual_heart', '{"shivada_jade", "prithiva_topaz"}'),
    ('pyro_hypostasis', 'Пиро гипостазис', 'inazuma', 'smoldering_pearl', '{"agnidus_agate"}'),
    ('hydro_hypostasis', 'Гидро гипостазис', 'inazuma', 'dew_of_repudiation', '{"varunada_lazurite"}'),
    ('thunder_manifestation', 'Манифестация грома', 'inazuma', 'storm_beads', '{"vajrada_amethyst"}'),
    ('golden_wolflord', 'Золотой волчий вожак', 'inazuma', 'riftborn_regalia', '{"prithiva_topaz"}'),
    ('coral_defenders', 'Стая вишапов глубин', 'inazuma', 'dragonheirs_false_fin', '{"shivada_jade", "vajrada_amethyst"}'),
    ('ruin_serpent', 'Змей руин', 'liyue', 'runic_fang', '{"prithiva_topaz"}')
;

-- +goose Down
DROP TABLE weekly_boss;
DROP TABLE world_boss;
