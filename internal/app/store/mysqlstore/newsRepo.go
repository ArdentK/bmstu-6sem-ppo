package mysqlstore

import "github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"

type NewsRepository struct {
	store *Store
}

func (r *NewsRepository) Create(*model.News) error       { return nil }
func (r *NewsRepository) Find(int) (*model.News, error)  { return nil, nil }
func (r *NewsRepository) GetAll() ([]*model.News, error) { return nil, nil }
