package storage

import "time"

// Interface задаёт контракт на работу с БД.
type Repository interface {
	Posts() ([]Post, error) // получение всех публикаций
	AddPost(Post) error     // создание новой публикации
	UpdatePost(Post) error  // обновление публикации
	DeletePost(Post) error  // удаление публикации по ID
}

// Storage is a storage.
type Storage struct {
	db *sqlx.DB
}

type Post struct {
	id         int64
	autorName  string
	header     string
	text       string
	createDate time.Time
}

func (s *Storage) Posts() ([]Post, error) {

	return nil, nil
} // получение всех публикаций

func (s *Storage) AddPost(Post) error {

	return nil
}

func (s *Storage) UpdatePost(Post) error {

	return nil
}

func (s *Storage) DeletePost(Post) error {

	return nil
}