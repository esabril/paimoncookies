-- +goose Up
CREATE TABLE IF NOT EXISTS weapon_material_type (
    id SERIAL PRIMARY KEY,
    title text NOT NULL,
    alias text,
    type text UNIQUE NOT NULL,
    location text NOT NULL
);

CREATE INDEX idx_weapon_material_type ON weapon_material_type (type);
CREATE INDEX idx_weapon_material_location ON weapon_material_type (location);

-- +goose Down
DROP TABLE weapon_material_type;

