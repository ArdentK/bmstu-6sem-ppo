package sqlstore

import (
	"database/sql"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
)

type AthletRepository struct {
	store *Store
}

func (r *AthletRepository) Create(a *model.Athlet) error {
	return r.store.db.QueryRow(
		"INSERT INTO battles (name, birthday, role_part, weapon_type, sex, rf_subject, sport_rank, sport_org) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		a.Name,
		a.Birthday,
		a.Role,
		a.WeaponType,
		a.Sex,
		a.RFSubject,
		a.Rank,
		a.SportOrg,
	).Scan(&a.ID)
}

func (r *AthletRepository) Find(id int) (*model.Athlet, error) {
	a := &model.Athlet{}
	err := r.store.db.QueryRow(
		"SELECT id, name, birthday, role_part, weapon_type, sex, rf_subject, sport_rank, sport_org FROM athlets WHERE id = $1",
		id,
	).Scan(
		&a.ID,
		&a.Name,
		&a.Birthday,
		&a.Role,
		&a.WeaponType,
		&a.Sex,
		&a.RFSubject,
		&a.Rank,
		&a.SportOrg,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return a, nil
}
func (r *AthletRepository) FindByWeaponType(weapon string) ([]*model.Athlet, error) { return nil, nil }
func (r *AthletRepository) FindBySex(sex int) ([]*model.Athlet, error)              { return nil, nil }
func (r *AthletRepository) FindByRFSubject(rfSubject string) ([]*model.Athlet, error) {
	return nil, nil
}
func (r *AthletRepository) FindByRank(rank string) ([]*model.Athlet, error) { return nil, nil }
