package teststore

import (
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
)

type BattleRepository struct {
	store   *Store
	battles map[int]*model.Battle
}

func (r *BattleRepository) Create(b *model.Battle) error                           { return nil }
func (r *BattleRepository) Find(id int) (*model.Battle, error)                     { return nil, nil }
func (r *BattleRepository) FindByIDWinner(idWinner int) ([]*model.Battle, error)   { return nil, nil }
func (r *BattleRepository) FindByIDLooser(idLooser int) ([]*model.Battle, error)   { return nil, nil }
func (r *BattleRepository) FindByIDReferee(idReferee int) ([]*model.Battle, error) { return nil, nil }
func (r *BattleRepository) FindByIDCompetition(idCompetition int) ([]*model.Battle, error) {
	return nil, nil
}
