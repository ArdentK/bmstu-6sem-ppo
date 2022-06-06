package store

import "github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

type CompetitionRepository interface {
	Create(*model.Competition) error
	GetAll() ([]*model.Competition, error)
	Find(int) (*model.Competition, error)
	Update(*model.Competition) error
	Delete(int) error
	FindByName(string) ([]*model.Competition, error)
	FindByDate(string) ([]*model.Competition, error)
	FindByAgeCategory(string) ([]*model.Competition, error)
	FindByWeaponType(string) ([]*model.Competition, error)
	FindByIsTeam(bool) ([]*model.Competition, error)
	FindByStatus(string) ([]*model.Competition, error)
	FindBySex(string) ([]*model.Competition, error)
	FindByType(string) ([]*model.Competition, error)
}

type BattleRepository interface {
	Create(*model.Battle) error
	Find(int) (*model.Battle, error)
	FindByIDWinner(int) ([]*model.Battle, error)
	FindByIDLooser(int) ([]*model.Battle, error)
	FindByIDReferee(int) ([]*model.Battle, error)
	FindByIDCompetition(int) ([]*model.Battle, error)
}

type AthletRepository interface {
	Create(*model.Athlet) error
	Find(int) (*model.Athlet, error)
	FindByWeaponType(string) ([]*model.Athlet, error)
	FindBySex(int) ([]*model.Athlet, error)
	FindByRFSubject(string) ([]*model.Athlet, error)
	FindByRank(string) ([]*model.Athlet, error)
}

type NewsRepository interface {
	Create(*model.News) error
	Find(int) (*model.News, error)
	GetAll() ([]*model.News, error)
}
