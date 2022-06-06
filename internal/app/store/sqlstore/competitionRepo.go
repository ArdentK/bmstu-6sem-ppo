package sqlstore

import (
	"database/sql"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
)

type CompetitionRepository struct {
	store *Store
}

func (r *CompetitionRepository) Create(c *model.Competition) error {
	return r.store.db.QueryRow(
		"INSERT INTO competitions (name, dt, age_category, weapon_type, is_team, status, sex, type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		c.Name,
		c.Date,
		c.AgeCategory,
		c.WeaponType,
		c.IsTeam,
		c.Status,
		c.Sex,
		c.Type,
	).Scan(&c.ID)
}
func (r *CompetitionRepository) Find(id int) (*model.Competition, error) {
	c := &model.Competition{}
	err := r.store.db.QueryRow(
		"SELECT id, name, dt, age_category, weapon_type, is_team, status, sex, type FROM competitions WHERE id = $1",
		id,
	).Scan(
		&c.ID,
		&c.Name,
		&c.Date,
		&c.AgeCategory,
		&c.WeaponType,
		&c.IsTeam,
		&c.Status,
		&c.Sex,
		&c.Type,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return c, nil
}

func (r *CompetitionRepository) GetAll() ([]*model.Competition, error) {
	items := []*model.Competition{}
	rows, err := r.store.db.Query("SELECT id, name, dt, age_category, weapon_type, is_team, status, sex, type FROM competitions;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		post := &model.Competition{}
		err = rows.Scan(
			&post.ID,
			&post.Name,
			&post.Date,
			&post.AgeCategory,
			&post.WeaponType,
			&post.IsTeam,
			&post.Status,
			&post.Sex,
			&post.Type,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, post)
	}

	return items, nil
}

func (r *CompetitionRepository) FindBySex(sex string) ([]*model.Competition, error) { return nil, nil }
func (r *CompetitionRepository) FindByType(t string) ([]*model.Competition, error)  { return nil, nil }
func (r *CompetitionRepository) FindByName(name string) ([]*model.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepository) FindByDate(dt string) ([]*model.Competition, error) { return nil, nil }
func (r *CompetitionRepository) FindByAgeCategory(ageCategory string) ([]*model.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepository) FindByWeaponType(weaponType string) ([]*model.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepository) FindByIsTeam(isTeam bool) ([]*model.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepository) FindByStatus(status string) ([]*model.Competition, error) {
	return nil, nil
}
