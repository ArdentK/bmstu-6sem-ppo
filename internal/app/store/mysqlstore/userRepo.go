package mysqlstore

import (
	"database/sql"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	res, err := r.store.db.Exec(
		"INSERT INTO users (email, encrypted_password) VALUES (?, ?)",
		u.Email,
		u.EncryptedPassword,
	)

	if err != nil {
		return err
	}

	newId, err := res.LastInsertId()
	u.ID = int(newId)

	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = ?",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = ?",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
