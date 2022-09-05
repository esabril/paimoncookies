-- +goose Up
-- +goose StatementBegin
ALTER TABLE talent_books_weekday
    DROP CONSTRAINT fk_talent_book_type,
    ADD CONSTRAINT fk_talent_book_type
        FOREIGN KEY (type)
            REFERENCES talent_books_type(type)
            ON DELETE CASCADE;

ALTER TABLE weapon_materials_weekday
    DROP CONSTRAINT fk_weapon_material_type,
    ADD CONSTRAINT fk_weapon_material_type
        FOREIGN KEY (type)
            REFERENCES weapon_materials_type(type)
            ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE talent_books_weekday
    DROP CONSTRAINT fk_talent_book_type,
    ADD CONSTRAINT fk_talent_book_type
        FOREIGN KEY (type)
            REFERENCES talent_books_type(type);

ALTER TABLE weapon_materials_weekday
    DROP CONSTRAINT fk_weapon_material_type,
    ADD CONSTRAINT fk_weapon_material_type
        FOREIGN KEY (type)
            REFERENCES weapon_materials_type(type);
-- +goose StatementEnd