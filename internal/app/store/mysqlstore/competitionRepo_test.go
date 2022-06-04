package mysqlstore_test

import (
	"testing"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store/mysqlstore"
	"github.com/stretchr/testify/assert"
)

func TestCompetitionRepo_Create(t *testing.T) {
	db, teardown := mysqlstore.TestDB(t, database, databaseURL)
	defer teardown("competitions")

	s := mysqlstore.New(db)
	c := model.TestCompetition(t)

	assert.NoError(t, s.Competition().Create(c))
	assert.NotNil(t, c)
}

func TestCompetitionRepo_Find(t *testing.T) {
	db, teardown := mysqlstore.TestDB(t, database, databaseURL)
	defer teardown("competitions")

	s := mysqlstore.New(db)
	c := model.TestCompetition(t)

	s.Competition().Create(c)

	c2, err := s.Competition().Find(c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, c2)
}
