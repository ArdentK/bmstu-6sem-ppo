package teststore

import (
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
)

type Store struct {
	userRepository        *UserRepository
	competitionRepository *CompetitionRepository
	battleRepository      *BattleRepository
	athletRepository      *AthletRepository
	newsRepository        *NewsRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) Competition() store.CompetitionRepository {
	if s.competitionRepository != nil {
		return s.competitionRepository
	}

	s.competitionRepository = &CompetitionRepository{
		store:        s,
		competitions: make(map[int]*model.Competition),
	}

	return s.competitionRepository
}

func (s *Store) Battle() store.BattleRepository {
	if s.battleRepository != nil {
		return s.battleRepository
	}

	s.battleRepository = &BattleRepository{
		store:   s,
		battles: make(map[int]*model.Battle),
	}

	return s.battleRepository
}

func (s *Store) Athlet() store.AthletRepository {
	if s.athletRepository != nil {
		return s.athletRepository
	}

	s.athletRepository = &AthletRepository{
		store:   s,
		athlets: make(map[int]*model.Athlet),
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
