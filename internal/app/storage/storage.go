package storage

import "errors"

type Storage interface {
	Add(key string, value string) error
	//Remove(key string)
	Get(key string) (string, error)
	Ping() error
	AddM(mas []URLRegistryM, short []string) error
}
type URLRegistryM struct {
	Id          string `json:"correlation_id"`
	OriginalUrl string `json:"original_url"`
}
type URLRegistryMRes struct {
	Id       string `json:"correlation_id"`
	ShortUrl string `json:"short_url"`
}

type MapStorage struct {
	m map[string]string
}

func NewMapStorage() Storage {
	return &MapStorage{
		m: make(map[string]string),
	}
}

func (s *MapStorage) Add(key string, value string) error { // я хуй знает как сюда ошибку запихнуть
	s.m[key] = value
	return nil
}

// func (s *MapStorage) Remove(key string) {
// 	delete(s.m, key)
// }

func (s *MapStorage) Get(key string) (string, error) {
	value, found := s.m[key]
	if !found {
		return "", errors.New("key not found")
	}
	return value, nil
}
func (s *MapStorage) Ping() error {
	return errors.New("тут нет бд")
}

func (s *MapStorage) AddM(mas []URLRegistryM, short []string) error {
	l := len(mas)
	for i := 0; i < l; i++ {
		s.Add(mas[i].OriginalUrl, short[i])
	}
	return nil
}
