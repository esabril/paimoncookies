-- +goose Up
INSERT INTO character (
    name, title, region, rarity, element, talent_book_type, talent_boss_drop,
    ascension_boss_drop, ascension_gem, ascension_local_speciality, common_ascension_material
) VALUES
    ('heizou', 'Сиканоин Хэйдзо', 'inazuma', 4, 'anemo', 'transience', 'the_meaning_of_aeons',
     'runic_fang', 'vayuda_turquoise', 'onikabuto', 'treasure_hoarder_insignias'),
    ('tighnari', 'Тигнари', 'sumeru',5 , 'dendro', 'admonition', 'the_meaning_of_aeons', 'majestic_hooked_beak',
     'nagadus_emerald', 'nilotpala_lotus', 'fungal_spore_powder'),
    ('collei', 'Коллеи', 'sumeru', 4, 'dendro', 'praxis', 'tears_of_the_calamitous_god',
     'majestic_hooked_beak', 'nagadus_emerald', 'rukkhashava_mushrooms', 'hilichurl_arrowheads')
;

-- +goose Down
DELETE FROM character WHERE name IN ('heizou', 'tighnari', 'collei');
