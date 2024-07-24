package storage

import "errors"

type Storage interface {
	Add(key string, value string, name string) error
	//Remove(key string)
	Get(key string) (string, error)
	Ping() error
	AddM(mas []URLRegistryM, short []string, name string) error
	Find(oru string) (string, error)
	ListUserURLs(Name string) ([]UserURL, error)
}
type UserURL struct {
	ShortURL    string
	OriginalURL string
}
type URLRegistryM struct {
	ID          string `json:"correlation_id"`
	OriginalURL string `json:"original_url"`
}
type URLRegistryMRes struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

type MapStorage struct {
	m map[string]UsersOriginal
}
type UsersOriginal struct {
	User     string
	Original string
}

func NewMapStorage() Storage {
	return &MapStorage{
		m: make(map[string]UsersOriginal),
	}
}

func (s *MapStorage) Add(key string, value string, name string) error { // я хуй знает как сюда ошибку запихнуть
	rez := UsersOriginal{User: name, Original: value}
	s.m[key] = rez
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
	return value.Original, nil
}
func (s *MapStorage) Ping() error {
	return errors.New("тут нет бд")
}

func (s *MapStorage) AddM(mas []URLRegistryM, short []string, name string) error {
	l := len(mas)
	for i := 0; i < l; i++ {
		s.Add(mas[i].OriginalURL, short[i], name)
	}
	return nil
}

func (s *MapStorage) Find(oru string) (string, error) {
	return "", errors.New("")
}

func (s *MapStorage) ListUserURLs(name string) ([]UserURL, error) {
	var rez []UserURL
	for k, z := range s.m {
		if z.User == name {
			rez = append(rez, UserURL{OriginalURL: z.Original, ShortURL: k})
		}
	}
	return rez, nil
}
