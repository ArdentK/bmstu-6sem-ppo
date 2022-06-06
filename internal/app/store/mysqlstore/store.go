package mysqlstore

import (
	"database/sql"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db                    *sql.DB
	userRepository        *UserRepository
	competitionRepository *CompetitionRepository
	battleRepository      *BattleRepository
	athletRepository      *AthletRepository
	newsRepository        *NewsRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Competition() store.CompetitionRepository {
	if s.competitionRepository != nil {
		return s.competitionRepository
	}

	s.competitionRepository = &CompetitionRepository{
		store: s,
	}

	return s.competitionRepository
}

func (s *Store) Battle() store.BattleRepository {
	if s.battleRepository != nil {
		return s.battleRepository
	}

	s.battleRepository = &BattleRepository{
		store: s,
	}

	return s.battleRepository
}

func (s *Store) Athlet() store.AthletRepository {
	if s.athletRepository != nil {
		return s.athletRepository
	}

	s.athletRepository = &AthletRepository{
		store: s,
	}

	return s.athletRepository
}

func (s *Store) News() store.NewsRepository {
	if s.newsRepository != nil {
		return s.newsRepository
	}

	s.newsRepository = &NewsRepository{
		store: s,
	}

	return s.newsRepository
}
