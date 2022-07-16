-- +goose Up
ALTER TABLE character
ADD COLUMN talent_boss_drop text,
ADD COLUMN ascension_boss_drop text,
ADD COLUMN ascension_gem text,
ADD COLUMN ascension_local_speciality text,
ADD COLUMN common_ascension_material text
;

UPDATE character
SET talent_boss_drop = 'dvalins_sigh',
    ascension_boss_drop = 'everflame_seed',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'small_lamp_grass',
    common_ascension_material = 'hilichurl_arrowheads'
WHERE name = 'amber';

UPDATE character
SET talent_boss_drop = 'spirit_locker_of_boreas',
    ascension_boss_drop = 'hoarfrost_core',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'calla_lily',
    common_ascension_material = 'treasure_hoarder_insignias'
WHERE name = 'kaeya';

UPDATE character
SET talent_boss_drop = 'dvalins_claw',
    ascension_boss_drop = 'lightning_prism',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'valberry',
    common_ascension_material = 'slime'
WHERE name = 'lisa';

UPDATE character
SET talent_boss_drop = 'ring_of_boreas',
    ascension_boss_drop = 'cleansing_heart',
    ascension_gem = 'varunada_lazurite',
    ascension_local_speciality = 'philanemo_mushroom',
    common_ascension_material = 'samachurl_scrolls'
WHERE name = 'barbara';

UPDATE character
SET talent_boss_drop = 'dvalins_claw',
    ascension_boss_drop = 'lightning_prism',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'wolfhook',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'raizor';

UPDATE character
SET talent_boss_drop = 'dvalins_claw',
    ascension_boss_drop = 'everflame_seed',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'jueyun_chili',
    common_ascension_material = 'slime'
WHERE name = 'xiangling';

UPDATE character
SET talent_boss_drop = 'dvalins_sigh',
    ascension_boss_drop = 'lightning_prism',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'noctilucous_jade',
    common_ascension_material = 'treasure_hoarder_insignias'
WHERE name = 'beidou';

UPDATE character
SET talent_boss_drop = 'tail_of_boreas',
    ascension_boss_drop = 'cleansing_heart',
    ascension_gem = 'varunada_lazurite',
    ascension_local_speciality = 'silk_flower',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'xingqui';

UPDATE character
SET talent_boss_drop = 'spirit_locker_of_boreas',
    ascension_boss_drop = 'basalt_pillar',
    ascension_gem = 'prithiva_topaz',
    ascension_local_speciality = 'glaze_lily',
    common_ascension_material = 'fatui_insignia'
WHERE name = 'ningguang';

UPDATE character
SET talent_boss_drop = 'spirit_locker_of_boreas',
    ascension_boss_drop = 'lightning_prism',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'small_lamp_grass',
    common_ascension_material = 'hilichurl_arrowheads'
WHERE name = 'fischl';

UPDATE character
SET talent_boss_drop = 'dvalins_plume',
    ascension_boss_drop = 'everflame_seed',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'windwheel_aster',
    common_ascension_material = 'treasure_hoarder_insignias'
WHERE name = 'bennett';

UPDATE character
SET talent_boss_drop = 'dvalins_claw',
    ascension_boss_drop = 'basalt_pillar',
    ascension_gem = 'prithiva_topaz',
    ascension_local_speciality = 'valberry',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'noelle';

UPDATE character
SET talent_boss_drop = 'dvalins_sigh',
    ascension_boss_drop = 'hoarfrost_core',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'cor_lapis',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'chongyun';

UPDATE character
SET talent_boss_drop = 'spirit_locker_of_boreas',
    ascension_boss_drop = 'hurricane_seed',
    ascension_gem = 'vayuda_turquoise',
    ascension_local_speciality = 'windwheel_aster',
    common_ascension_material = 'whopperflower_nectar'
WHERE name = 'sucrose';

UPDATE character
SET talent_boss_drop = 'dvalins_plume',
    ascension_boss_drop = 'hurricane_seed',
    ascension_gem = 'vayuda_turquoise',
    ascension_local_speciality = 'dandelion_seed',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'jean';

UPDATE character
SET talent_boss_drop = 'dvalins_plume',
    ascension_boss_drop = 'everflame_seed',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'small_lamp_grass',
    common_ascension_material = 'fatui_insignia'
WHERE name = 'diluc';

UPDATE character
SET talent_boss_drop = 'tail_of_boreas',
    ascension_boss_drop = 'hoarfrost_core',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'violetgrass',
    common_ascension_material = 'samachurl_scrolls'
WHERE name = 'qiqi';

UPDATE character
SET talent_boss_drop = 'ring_of_boreas',
    ascension_boss_drop = 'cleansing_heart',
    ascension_gem = 'varunada_lazurite',
    ascension_local_speciality = 'philanemo_mushroom',
    common_ascension_material = 'whopperflower_nectar'
WHERE name = 'mona';

UPDATE character
SET talent_boss_drop = 'ring_of_boreas',
    ascension_boss_drop = 'lightning_prism',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'cor_lapis',
    common_ascension_material = 'whopperflower_nectar'
WHERE name = 'keqing';

UPDATE character
SET talent_boss_drop = 'tail_of_boreas',
    ascension_boss_drop = 'hurricane_seed',
    ascension_gem = 'vayuda_turquoise',
    ascension_local_speciality = 'cecilia',
    common_ascension_material = 'slime'
WHERE name = 'venti';

UPDATE character
SET talent_boss_drop = 'ring_of_boreas',
    ascension_boss_drop = 'everflame_seed',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'philanemo_mushroom',
    common_ascension_material = 'samachurl_scrolls'
WHERE name = 'klee';

UPDATE character
SET talent_boss_drop = 'shard_of_foul_legacy',
    ascension_boss_drop = 'hoarfrost_core',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'calla_lily',
    common_ascension_material = 'hilichurl_arrowheads'
WHERE name = 'diona';

UPDATE character
SET talent_boss_drop = 'shard_of_foul_legacy',
    ascension_boss_drop = 'cleansing_heart',
    ascension_gem = 'varunada_lazurite',
    ascension_local_speciality = 'starconch',
    common_ascension_material = 'fatui_insignia'
WHERE name = 'tartaglia';

UPDATE character
SET talent_boss_drop = 'tusk_of_monoceros_caeli',
    ascension_boss_drop = 'everflame_seed',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'violetgrass',
    common_ascension_material = 'treasure_hoarder_insignias'
WHERE name = 'xinyan';

UPDATE character
SET talent_boss_drop = 'tusk_of_monoceros_caeli',
    ascension_boss_drop = 'basalt_pillar',
    ascension_gem = 'prithiva_topaz',
    ascension_local_speciality = 'cor_lapis',
    common_ascension_material = 'slime'
WHERE name = 'zhongli';

UPDATE character
SET talent_boss_drop = 'tusk_of_monoceros_caeli',
    ascension_boss_drop = 'basalt_pillar',
    ascension_gem = 'prithiva_topaz',
    ascension_local_speciality = 'cecilia',
    common_ascension_material = 'samachurl_scrolls'
WHERE name = 'albedo';

UPDATE character
SET talent_boss_drop = 'shadow_of_the_warrior',
    ascension_boss_drop = 'hoarfrost_core',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'qingxin',
    common_ascension_material = 'whopperflower_nectar'
WHERE name = 'ganyu';

UPDATE character
SET talent_boss_drop = 'shadow_of_the_warrior',
    ascension_boss_drop = 'juvenile_jade',
    ascension_gem = 'vayuda_turquoise',
    ascension_local_speciality = 'qingxin',
    common_ascension_material = 'slime'
WHERE name = 'xiao';

UPDATE character
SET talent_boss_drop = 'shard_of_foul_legacy',
    ascension_boss_drop = 'juvenile_jade',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'silk_flower',
    common_ascension_material = 'whopperflower_nectar'
WHERE name = 'hutao';

UPDATE character
SET talent_boss_drop = 'shadow_of_the_warrior',
    ascension_boss_drop = 'hoarfrost_core',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'valberry',
    common_ascension_material = 'fatui_insignia'
WHERE name = 'rosaria';

UPDATE character
SET talent_boss_drop = 'bloodjade_branch',
    ascension_boss_drop = 'juvenile_jade',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'noctilucous_jade',
    common_ascension_material = 'treasure_hoarder_insignias'
WHERE name = 'yanfei';

UPDATE character
SET talent_boss_drop = 'dragon_lords_crown',
    ascension_boss_drop = 'crystalline_bloom',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'dandelion_seed',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'eula';

UPDATE character
SET talent_boss_drop = 'gilded_scale',
    ascension_boss_drop = 'marionette_core',
    ascension_gem = 'vayuda_turquoise',
    ascension_local_speciality = 'sea_ganoderma',
    common_ascension_material = 'treasure_hoarder_insignias'
WHERE name = 'kadzukha';

UPDATE character
SET talent_boss_drop = 'bloodjade_branch',
    ascension_boss_drop = 'perpetual_heart',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'sakura_bloom',
    common_ascension_material = 'nobushi_handguards'
WHERE name = 'ayaka';

UPDATE character
SET talent_boss_drop = 'gilded_scale',
    ascension_boss_drop = 'marionette_core',
    ascension_gem = 'vayuda_turquoise',
    ascension_local_speciality = 'crystal_marrow',
    common_ascension_material = 'whopperflower_nectar'
WHERE name = 'sayu';

UPDATE character
SET talent_boss_drop = 'dragon_lords_crown',
    ascension_boss_drop = 'smoldering_pearl',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'naku_weed',
    common_ascension_material = 'samachurl_scrolls'
WHERE name = 'yoimiya';

UPDATE character
SET talent_boss_drop = 'ashen_heart',
    ascension_boss_drop = 'storm_beads',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'dendrobium',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'sara';

UPDATE character
SET talent_boss_drop = 'molten_moment',
    ascension_boss_drop = 'crystalline_bloom',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'crystal_marrow',
    common_ascension_material = 'spectral_cores'
WHERE name = 'aloy';

UPDATE character
SET talent_boss_drop = 'molten_moment',
    ascension_boss_drop = 'storm_beads',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'amakumo_fruit',
    common_ascension_material = 'nobushi_handguards'
WHERE name = 'raiden';

UPDATE character
SET talent_boss_drop = 'hellfire_butterfly',
    ascension_boss_drop = 'dew_of_repudiation',
    ascension_gem = 'varunada_lazurite',
    ascension_local_speciality = 'sango_pearl',
    common_ascension_material = 'spectral_cores'
WHERE name = 'kokomi';

UPDATE character
SET talent_boss_drop = 'hellfire_butterfly',
    ascension_boss_drop = 'smoldering_pearl',
    ascension_gem = 'agnidus_agate',
    ascension_local_speciality = 'fluorescent_fungus',
    common_ascension_material = 'treasure_hoarder_insignias'
WHERE name = 'toma';

UPDATE character
SET talent_boss_drop = 'molten_moment',
    ascension_boss_drop = 'perpetual_heart',
    ascension_gem = 'prithiva_topaz',
    ascension_local_speciality = 'sango_pearl',
    common_ascension_material = 'spectral_cores'
WHERE name = 'gorou';

UPDATE character
SET talent_boss_drop = 'ashen_heart',
    ascension_boss_drop = 'riftborn_regalia',
    ascension_gem = 'prithiva_topaz',
    ascension_local_speciality = 'onikabuto',
    common_ascension_material = 'slime'
WHERE name = 'itto';

UPDATE character
SET talent_boss_drop = 'ashen_heart',
    ascension_boss_drop = 'riftborn_regalia',
    ascension_gem = 'prithiva_topaz',
    ascension_local_speciality = 'glaze_lily',
    common_ascension_material = 'hilichurl_masks'
WHERE name = 'yunjin';

UPDATE character
SET talent_boss_drop = 'hellfire_butterfly',
    ascension_boss_drop = 'dragonheirs_false_fin',
    ascension_gem = 'shivada_jade',
    ascension_local_speciality = 'qingxin',
    common_ascension_material = 'whopperflower_nectar'
WHERE name = 'shenhe';

UPDATE character
SET talent_boss_drop = 'the_meaning_of_aeons',
    ascension_boss_drop = 'dragonheirs_false_fin',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'sea_ganoderma',
    common_ascension_material = 'nobushi_handguards'
WHERE name = 'yaemiko';

UPDATE character
SET talent_boss_drop = 'mudra_of_the_malefic_general',
    ascension_boss_drop = 'dew_of_repudiation',
    ascension_gem = 'varunada_lazurite',
    ascension_local_speciality = 'sakura_bloom',
    common_ascension_material = 'nobushi_handguards'
WHERE name = 'ayato';

UPDATE character
SET talent_boss_drop = '',
    ascension_boss_drop = '',
    ascension_gem = '',
    ascension_local_speciality = '',
    common_ascension_material = ''
WHERE name = '';

UPDATE character
SET talent_boss_drop = 'gilded_scale',
    ascension_boss_drop = 'runic_fang',
    ascension_gem = 'varunada_lazurite',
    ascension_local_speciality = 'starconch',
    common_ascension_material = 'fatui_insignia'
WHERE name = 'yelan';

UPDATE character
SET talent_boss_drop = 'tears_of_the_calamitous_god',
    ascension_boss_drop = 'runic_fang',
    ascension_gem = 'vajrada_amethyst',
    ascension_local_speciality = 'naku_weed',
    common_ascension_material = 'spectral_cores'
WHERE name = 'kuki';

-- +goose Down
ALTER TABLE character
DROP COLUMN talent_boss_drop,
DROP COLUMN ascension_boss_drop,
DROP COLUMN ascension_gem,
DROP COLUMN ascension_local_speciality,
DROP COLUMN common_ascension_material
;