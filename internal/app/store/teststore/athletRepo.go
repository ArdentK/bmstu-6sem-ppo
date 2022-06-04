package teststore

import "github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"

type AthletRepository struct {
	store   *Store
	athlets map[int]*model.Athlet
}

func (r *AthletRepository) Create(a *model.Athlet) error                            { return nil }
func (r *AthletRepository) Find(int) (*model.Athlet, error)                         { return nil, nil }
func (r *AthletRepository) FindByWeaponType(weapon string) ([]*model.Athlet, error) { return nil, nil }
func (r *AthletRepository) FindBySex(sex int) ([]*model.Athlet, error)              { return nil, nil }
func (r *AthletRepository) FindByRFSubject(rfSubject string) ([]*model.Athlet, error) {
	return nil, nil
}
func (r *AthletRepository) FindByRank(rank string) ([]*model.Athlet, error) { return nil, nil }
