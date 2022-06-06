package store

type Store interface {
	User() UserRepository
	Competition() CompetitionRepository
	Battle() BattleRepository
	Athlet() AthletRepository
	News() NewsRepository
}
