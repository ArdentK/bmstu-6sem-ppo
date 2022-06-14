package mysqlstore_test

import (
	"testing"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store/mysqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := mysqlstore.TestDB(t, database, databaseURL)
	defer teardown("users")

	s := mysqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := mysqlstore.TestDB(t, database, databaseURL)
	defer teardown("users")

	s := mysqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Create(t *testing.T) {
	db, teardown := mysqlstore.TestDB(t, database, databaseURL)
	defer teardown("users")

	s := mysqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}
