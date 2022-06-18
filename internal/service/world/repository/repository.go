package repository

import (
	"errors"
	"github.com/esabril/paimoncookies/internal/service/world/model"
	"github.com/jmoiron/sqlx"
)

// IWorldRepo common repo interface
type IWorldRepo interface {
	GetWeekdayTalentBooksWithLocation(weekday string) ([]model.TalentBook, error)
	GetWeekdayWeaponMaterialsWithLocation(weekday string) ([]model.WeaponMaterial, error)
	GetRegions() ([]model.Region, error)
	GetTalentBookByType(bookType string) (model.TalentBook, error)
	GetTalentBookWeekdays(bookType string) ([]string, error)
}

// Repository
type repo struct {
	db *sqlx.DB
}

// New initialize new repository
func New(db *sqlx.DB) IWorldRepo {
	return &repo{
		db: db,
	}
}

// GetWeekdayTalentBooksWithLocation returns list of today's talent books in dungeons
func (r *repo) GetWeekdayTalentBooksWithLocation(weekday string) ([]model.TalentBook, error) {
	// books := make(map[string][]string)
	books := make([]model.TalentBook, 0)
	args := map[string]interface{}{
		"weekday": weekday,
	}

	query := `SELECT r.title AS location, t.title AS title FROM talent_books_type AS t
			JOIN talent_books_weekday AS tbw ON t.type = tbw.type
			JOIN region r on t.location = r.name
			WHERE weekday = :weekday
			ORDER BY r.id`

	rows, err := r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b model.TalentBook

		err = rows.StructScan(&b)
		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}

// GetWeekdayWeaponMaterialsWithLocation returns list of today's weapon materials in dungeons
func (r *repo) GetWeekdayWeaponMaterialsWithLocation(weekday string) ([]model.WeaponMaterial, error) {
	materials := make([]model.WeaponMaterial, 0)
	args := map[string]interface{}{
		"weekday": weekday,
	}

	query := `SELECT r.title AS location, t.title AS title, t.alias AS alias FROM weapon_materials_type AS t
    		JOIN weapon_materials_weekday AS wmw ON t.type = wmw.type
    		JOIN region r on t.location = r.name
			WHERE weekday = :weekday
			ORDER BY r.id`

	rows, err := r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var m model.WeaponMaterial

		err = rows.StructScan(&m)
		if err != nil {
			return nil, err
		}

		materials = append(materials, m)
	}

	return materials, nil
}

// GetRegions in strict order
func (r *repo) GetRegions() ([]model.Region, error) {
	var regions []model.Region
	query := "SELECT id, name, title FROM region ORDER BY id"

	err := r.db.Select(&regions, query)

	return regions, err
}

func (r *repo) GetTalentBookByType(bookType string) (model.TalentBook, error) {
	var books []model.TalentBook
	args := map[string]interface{}{
		"bookType": bookType,
	}

	stmt, err := r.db.PrepareNamed(`SELECT bt.id, bt.title, bt.type, r.title as location 
		FROM talent_books_type bt 
		    JOIN region r on bt.location = r.name
        WHERE type = :bookType LIMIT 1`,
	)
	if err != nil {
		return model.TalentBook{}, err
	}

	err = stmt.Select(&books, args)
	if err != nil {
		return model.TalentBook{}, err
	}

	if len(books) == 0 {
		return model.TalentBook{}, errors.New("Talent Book not found")
	}

	return books[0], err
}

func (r *repo) GetTalentBookWeekdays(bookType string) ([]string, error) {
	books := make([]string, 0)
	args := map[string]interface{}{
		"type": bookType,
	}

	stmt, err := r.db.PrepareNamed(`SELECT weekday FROM talent_books_weekday WHERE type = :type`)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&books, args)

	return books, err
}
