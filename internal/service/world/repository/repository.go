package repository

import (
	"github.com/esabril/paimoncookies/internal/service/world/model"
	"github.com/jmoiron/sqlx"
)

// IRepo common repo interface
type IRepo interface {
	GetWeekdayTalentBooksWithLocation(weekday string) ([]model.TalentBook, error)
	GetWeekdayWeaponMaterialsWithLocation(weekday string) ([]model.WeaponMaterial, error)
}

// Repository
type repo struct {
	db *sqlx.DB
}

// New initialize new repository
func New(db *sqlx.DB) IRepo {
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
