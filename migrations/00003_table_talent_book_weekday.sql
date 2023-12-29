-- +goose Up
CREATE TABLE IF NOT EXISTS talent_book_weekday (
    type text,
    weekday text,

    CONSTRAINT fk_talent_book_type
        FOREIGN KEY (type)
            REFERENCES talent_book_type(type)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);

CREATE INDEX idx_weekday ON talent_book_weekday (weekday);

-- +goose Down
DROP TABLE talent_book_weekday;
