package sqlstore

import (
	"database/sql"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
)

type NewsRepository struct {
	store *Store
}

func (r *NewsRepository) Create(n *model.News) error {
	return r.store.db.QueryRow(
		"INSERT INTO news (title, description) VALUES ($1, $2) RETURNING id",
		n.Title,
		n.Description,
	).Scan(&n.ID)
}
func (r *NewsRepository) Find(id int) (*model.News, error) {
	n := &model.News{}
	err := r.store.db.QueryRow(
		"SELECT title, description FROM news WHERE id = $1",
		id,
	).Scan(
		&n.ID,
		&n.Title,
		&n.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return n, nil
}
func (r *NewsRepository) GetAll() ([]*model.News, error) {
	items := []*model.News{}
	rows, err := r.store.db.Query("SELECT id, title, description FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &model.News{}
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, post)
	}

	return items, nil
}
