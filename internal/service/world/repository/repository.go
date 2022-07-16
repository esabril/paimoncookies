package repository

import (
	"errors"
	"fmt"
	"github.com/esabril/paimoncookies/internal/service/world/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

// IWorldRepo common repo interface
type IWorldRepo interface {
	GetWeekdayTalentBooksWithLocation(weekday string) ([]model.TalentBook, error)
	GetWeekdayWeaponMaterialsWithLocation(weekday string) ([]model.WeaponMaterial, error)
	GetRegions() ([]model.Region, error)
	GetTalentBookByType(bookType string) (model.TalentBook, error)
	GetTalentBookWeekdays(bookType string) ([]string, error)
	GetGemByName(name string) (model.Gem, error)
	GetWeeklyBossDropByName(name string) (model.BossDrop, error)
	GetWorldBossDropByName(name string) (model.BossDrop, error)
	GetAscensionMaterialsByNames(names []string) ([]model.AscensionMaterial, error)
	FindGemByTitle(title string) (model.Gem, error)
	GetGemDropInfoByName(name string) ([]model.BossDrop, error)
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
		return model.TalentBook{}, errors.New("talent Book not found")
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

func (r *repo) GetGemByName(name string) (model.Gem, error) {
	var g []model.Gem
	args := map[string]interface{}{
		"name": name,
	}

	stmt, err := r.db.PrepareNamed(`SELECT * FROM gem WHERE name = :name LIMIT 1`)
	if err != nil {
		return model.Gem{}, err
	}

	err = stmt.Select(&g, args)
	if err != nil {
		return model.Gem{}, err
	}

	if len(g) == 0 {
		return model.Gem{}, errors.New("gem not found")
	}

	return g[0], nil
}

func (r *repo) GetWeeklyBossDropByName(name string) (model.BossDrop, error) {
	var bd []model.BossDrop

	args := map[string]interface{}{
		"name": name,
	}

	stmt, err := r.db.PrepareNamed(
		`SELECT bd.title as title, wb.title as boss, wb.domain as domain, r.title AS location, 'weekly' AS type  
				FROM boss_drop AS bd
					JOIN weekly_boss AS wb ON bd.name = ANY(wb.talent_materials)
					JOIN region AS r ON wb.location = r.name
				WHERE bd.name = :name LIMIT 1`,
	)
	if err != nil {
		return model.BossDrop{}, err
	}

	err = stmt.Select(&bd, args)
	if err != nil {
		return model.BossDrop{}, err
	}

	if len(bd) == 0 {
		return model.BossDrop{}, errors.New("weekly boss drop not found")
	}

	return bd[0], nil
}

func (r *repo) GetWorldBossDropByName(name string) (model.BossDrop, error) {
	var bd []model.BossDrop

	args := map[string]interface{}{
		"name": name,
	}

	stmt, err := r.db.PrepareNamed(
		`SELECT bd.title as title, wb.title as boss, r.title AS location, 'world' AS type  
				FROM boss_drop AS bd
					JOIN world_boss AS wb ON bd.name = wb.ascension_material
					JOIN region AS r ON wb.location = r.name
				WHERE bd.name = :name LIMIT 1`,
	)

	if err != nil {
		return model.BossDrop{}, err
	}

	err = stmt.Select(&bd, args)
	if err != nil {
		return model.BossDrop{}, err
	}

	if len(bd) == 0 {
		return model.BossDrop{}, errors.New("world boss drop not found")
	}

	return bd[0], nil
}

func (r *repo) GetAscensionMaterialsByNames(names []string) ([]model.AscensionMaterial, error) {
	result := make([]model.AscensionMaterial, 0)

	arg := map[string]interface{}{
		"names": names,
	}

	query, args, err := sqlx.Named(`SELECT title, type FROM ascension_material WHERE name IN (:names)`, arg)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var m model.AscensionMaterial

		err = rows.StructScan(&m)
		if err != nil {
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
}

func (r *repo) FindGemByTitle(title string) (model.Gem, error) {
	var g []model.Gem
	title = strings.ToLower(title)

	args := map[string]interface{}{
		"title": fmt.Sprintf("%%%s%%", title),
	}

	stmt, err := r.db.PrepareNamed(`SELECT * FROM gem WHERE LOWER(title) LIKE :title LIMIT 1`)
	if err != nil {
		return model.Gem{}, err
	}

	err = stmt.Select(&g, args)
	if err != nil {
		return model.Gem{}, err
	}

	if len(g) == 0 {
		return model.Gem{}, errors.New("gem not found")
	}

	return g[0], err
}

func (r *repo) GetGemDropInfoByName(name string) ([]model.BossDrop, error) {
	result := make([]model.BossDrop, 0)

	args := map[string]interface{}{
		"name": name,
	}

	stmt, err := r.db.PrepareNamed(
		`SELECT wkb.id     AS id,
			   wkb.title  AS boss,
			   r.title    AS location,
			   wkb.domain AS domain,
			   'weekly'   AS type
		FROM weekly_boss AS wkb
				 JOIN region r on wkb.location = r.name
		WHERE :name = ANY (wkb.gems)
		UNION ALL
		SELECT wb.id    AS id,
			   wb.title AS boss,
			   r.title  AS location,
			   ''       AS domain,
			   'world'  AS type
		FROM world_boss AS wb
				 JOIN region r on wb.location = r.name
		WHERE :name = ANY (wb.gems)
		ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&result, args)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("there is no information about gem's drop info")
	}

	return result, nil
}
