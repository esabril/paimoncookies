-- +goose Up
CREATE TABLE IF NOT EXISTS boss_drop (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text,
    type text
);

CREATE INDEX idx_character_boss_drop_type ON boss_drop (
    type
);

INSERT INTO boss_drop (name, title, type) VALUES
   ('hurricane_seed', 'Семя урагана', 'ascension'),
   ('lightning_prism', 'Призма молнии', 'ascension'),
   ('basalt_pillar', 'Базальтовая колонна', 'ascension'),
   ('hoarfrost_core', 'Инеевое ядро', 'ascension'),
   ('everflame_seed', 'Пылающее семя', 'ascension'),
   ('cleansing_heart', 'Очищающее сердце', 'ascension'),
   ('juvenile_jade', 'Незрелый нефрит', 'ascension'),
   ('crystalline_bloom', 'Цветение кристалла', 'ascension'),
   ('marionette_core', 'Ядро марионетки', 'ascension'),
   ('perpetual_heart', 'Сердце бесконечного механизма', 'ascension'),
   ('smoldering_pearl', 'Тлеющая жемчужина', 'ascension'),
   ('dew_of_repudiation', 'Роса отречения', 'ascension'),
   ('storm_beads', 'Штормовой жемчуг', 'ascension'),
   ('riftborn_regalia', 'Эмблема Разрыва', 'ascension'),
   ('dragonheirs_false_fin', 'Ложный плавник Дракона глубин', 'ascension'),
   ('runic_fang', 'Рунический клык', 'ascension'),
   ('dvalins_plume', 'Перо из хвоста Двалина', 'talent'),
   ('dvalins_claw', 'Коготь Двалина', 'talent'),
   ('dvalins_sigh', 'Вздох Двалина', 'talent'),
   ('tail_of_boreas', 'Хвост Борея', 'talent'),
   ('ring_of_boreas', 'Кольцо Борея', 'talent'),
   ('spirit_locker_of_boreas', 'Шкатулка с духом Борея', 'talent'),
   ('tusk_of_monoceros_caeli', 'Рог небесного кита', 'talent'),
   ('shard_of_foul_legacy', 'Осколок дьявольского меча', 'talent'),
   ('shadow_of_the_warrior', 'Тень воина', 'talent'),
   ('dragon_lords_crown', 'Корона лорда драконов', 'talent'),
   ('bloodjade_branch', 'Ветвь кровавой яшмы', 'talent'),
   ('gilded_scale', 'Позолоченная чешуя', 'talent'),
   ('molten_moment', 'Расплавленный миг', 'talent'),
   ('hellfire_butterfly', 'Бабочка адского пламени', 'talent'),
   ('ashen_heart', 'Пепельное сердце', 'talent'),
   ('mudra_of_the_malefic_general', 'Мудра зловещего генерала', 'talent'),
   ('tears_of_the_calamitous_god', 'Слёзы очищения божества бедствий', 'talent'),
   ('the_meaning_of_aeons', 'Смысл эонов', 'talent')
;
-- +goose Down
DROP TABLE boss_drop;