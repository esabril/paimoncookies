-- +goose Up
CREATE TABLE IF NOT EXISTS talent_books_weekday (
    type text,
    weekday text,
    CONSTRAINT fk_talent_book_type
        FOREIGN KEY (type)
            REFERENCES talent_books_type(type)
);

CREATE INDEX idx_weekday ON talent_books_weekday (
    weekday
);

INSERT INTO talent_books_weekday (type, weekday) VALUES
     ('freedom', 'monday'),
     ('freedom', 'thursday'),
     ('prosperity', 'monday'),
     ('prosperity', 'thursday'),
     ('transience', 'monday'),
     ('transience', 'thursday'),
     ('resistance', 'tuesday'),
     ('resistance', 'friday'),
     ('diligence', 'tuesday'),
     ('diligence', 'friday'),
     ('elegance', 'tuesday'),
     ('elegance', 'friday'),
     ('ballad', 'wednesday'),
     ('ballad', 'saturday'),
     ('gold', 'wednesday'),
     ('gold', 'saturday'),
     ('light', 'wednesday'),
     ('light', 'saturday')
;

-- +goose Down
DROP TABLE talent_books_weekday;
