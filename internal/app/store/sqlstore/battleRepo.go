package sqlstore

import (
	"database/sql"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
)

type BattleRepository struct {
	store *Store
}

func (r *BattleRepository) Create(b *model.Battle) error {
	return r.store.db.QueryRow(
		"INSERT INTO battles (id_winner, id_looser, id_competiton, winner_score, looser_score) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		b.IDWinner,
		b.IDLooser,
		b.IDCompetition,
		b.WinnerScore,
		b.LooserScore,
	).Scan(&b.ID)
}

func (r *BattleRepository) Find(id int) (*model.Battle, error) {
	b := &model.Battle{}
	err := r.store.db.QueryRow(
		"SELECT id, id_winner, id_looser, id_competition, winner_score, looser_score FROM battles WHERE id = $1",
		id,
	).Scan(
		&b.ID,
		&b.IDWinner,
		&b.IDLooser,
		&b.IDCompetition,
		&b.WinnerScore,
		&b.LooserScore,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return b, nil
}
func (r *BattleRepository) FindByIDWinner(idWinner int) ([]*model.Battle, error)   { return nil, nil }
func (r *BattleRepository) FindByIDLooser(idLooser int) ([]*model.Battle, error)   { return nil, nil }
func (r *BattleRepository) FindByIDReferee(idReferee int) ([]*model.Battle, error) { return nil, nil }
func (r *BattleRepository) FindByIDCompetition(idCompetition int) ([]*model.Battle, error) {
	return nil, nil
}
