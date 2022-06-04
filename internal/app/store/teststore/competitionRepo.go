package teststore

import "github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"

type CompetitionRepository struct {
	store        *Store
	competitions map[int]*model.Competition
}

func (r *CompetitionRepository) Create(c *model.Competition) error       { return nil }
func (r *CompetitionRepository) GetAll() ([]*model.Competition, error)   { return nil, nil }
func (r *CompetitionRepository) Find(id int) (*model.Competition, error) { return nil, nil }
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
func (r *CompetitionRepository) FindBySex(sex string) ([]*model.Competition, error) { return nil, nil }
func (r *CompetitionRepository) FindByType(t string) ([]*model.Competition, error)  { return nil, nil }
