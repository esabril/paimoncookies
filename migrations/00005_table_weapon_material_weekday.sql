-- +goose Up
CREATE TABLE IF NOT EXISTS weapon_material_weekday (
    type text,
    weekday text,

    CONSTRAINT fk_weapon_material_type
        FOREIGN KEY (type)
            REFERENCES weapon_material_type(type)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE weapon_material_weekday;

