-- +goose Up
CREATE TABLE IF NOT EXISTS character(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    title text NOT NULL,
    region text NOT NULL,
    rarity int8 NOT NULL,
    element text NOT NULL,
    talent_book_type text
);

CREATE INDEX idx_character_name ON character (title);
CREATE INDEX idx_character_element ON character (element);
CREATE INDEX idx_character_talent_book_type ON character (talent_book_type);

INSERT INTO character (name, title, region, rarity, element, talent_book_type) VALUES
    ('amber', 'Эмбер', 'mondstadt', 4, 'pyro', 'freedom'),
    ('kaeya', 'Кэйа', 'mondstadt', 4, 'cryo', 'ballad'),
    ('lisa', 'Лиза', 'mondstadt', 4, 'electro', 'ballad'),
    ('barbara', 'Барбара', 'mondstadt', 4, 'hydro', 'freedom'),
    ('raizor', 'Рэйзор', 'mondstadt', 4, 'electro', 'resistance'),
    ('xiangling', 'Сян Лин', 'liyue', 4, 'pyro', 'diligence'),
    ('beidou', 'Бэй Доу', 'liyue', 4, 'electro', 'gold'),
    ('xingqui', 'Син Цю', 'liyue', 4, 'hydro', 'gold'),
    ('ningguang', 'Нин Гуан', 'liyue', 4, 'geo', 'prosperity'),
    ('fischl', 'Фишль', 'mondstadt', 4, 'electro', 'ballad'),
    ('bennet', 'Беннет', 'mondstadt', 4, 'pyro', 'resistance'),
    ('noelle', 'Ноэлль', 'mondstadt', 4, 'geo', 'resistance'),
    ('chongyun', 'Чунь Юнь', 'liyue', 4, 'cryo', 'diligence'),
    ('sucrose', 'Сахароза', 'mondstadt', 4, 'anemo', 'freedom'),
    ('jean', 'Джин', 'mondstadt', 5, 'anemo', 'resistance'),
    ('diluc', 'Дилюк', 'mondstadt', 5, 'pyro', 'resistance'),
    ('qiqi', 'Ци Ци', 'liyue', 5, 'cryo', 'prosperity'),
    ('mona', 'Мона', 'mondstadt', 5, 'hydro', 'resistance'),
    ('keqing', 'Кэ Цин', 'liyue', 5, 'electro', 'prosperity'),
    ('venti', 'Венти', 'mondstadt', 5, 'anemo', 'ballad'),
    ('klee', 'Кли', 'mondstadt', 5, 'pyro', 'freedom'),
    ('diona', 'Диона', 'mondstadt', 4, 'cryo', 'freedom'),
    ('tartaglia', 'Тарталья', 'liyue', 5, 'hydro', 'freedom'),
    ('xinyan', 'Синь Янь', 'liyue', 4, 'pyro', 'gold'),
    ('zhongli', 'Чжун Ли', 'liyue', 5, 'geo', 'gold'),
    ('albedo', 'Альбедо', 'mondstadt', 5, 'geo', 'ballad'),
    ('ganyu', 'Гань Юй', 'liyue', 5, 'cryo', 'diligence'),
    ('xiao', 'Сяо', 'liyue', 5, 'anemo', 'prosperity'),
    ('hutao', 'Ху Тао', 'liyue', 5, 'pyro', 'diligence'),
    ('rosaria', 'Розария', 'mondstadt', 4, 'cryo', 'ballad'),
    ('yanfei', 'Янь Фэй', 'liyue', 4, 'pyro', 'gold'),
    ('eola', 'Эола', 'mondstadt', 5, 'cryo', 'resistance'),
    ('kadzukha', 'Каэдэхара Кадзуха', 'inazuma', 5, 'anemo', 'diligence'),
    ('ayaka', 'Камисато Аяка', 'inazuma', 5, 'cryo', 'elegance'),
    ('sayu', 'Саю', 'inazuma', 4, 'anemo', 'light'),
    ('yoimiya', 'Йоимия', 'inazuma', 5, 'pyro', 'transience'),
    ('sara', 'Кудзё Сара', 'inazuma', 4, 'electro', 'elegance'),
    ('eloy', 'Элой', 'Вселенная «Horizon Zero Dawn»', 0, 'cryo', 'freedom'),
    ('rayden', 'Рейден', 'inazuma', 5, 'electro', 'light'),
    ('kokomi', 'Сангономия Кокоми', 'inazuma', 5, 'hydro', 'transience'),
    ('toma', 'Тома', 'inazuma', 4, 'pyro', 'transience'),
    ('gorou', 'Горо', 'inazuma', 4, 'geo', 'light'),
    ('itto', 'Аратаки Итто', 'inazuma', 5, 'geo', 'elegance'),
    ('yunjin', 'Юнь Цзинь', 'liyue', 4, 'geo', 'diligence'),
    ('shenhe', 'Шэнь Хэ', 'liyue', 5, 'cryo', 'prosperity'),
    ('yaemiko', 'Яэ Мико', 'inazuma', 5, 'electro', 'light'),
    ('ayato', 'Камисато Аято', 'inazuma', 5, 'hydro', 'elegance'),
    ('yelan', 'Е Лань', 'liyue', 5, 'hydro', 'prosperity'),
    ('kuki', 'Куки Синобу', 'inazuma', 4, 'electro', 'elegance')
;

-- +goose Down
DROP TABLE character;
