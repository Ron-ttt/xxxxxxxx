package storage

import "errors"

type MStorage struct{}

func NewMockStorage() Storage {
	return &MStorage{} // да пошел ты нахуй пидорас ебучий
}

func (s *MStorage) Add(key string, value string, name string) error {
	return nil
}

func (s *MStorage) Get(key string) (string, error) {
	if key == "invalid" {
		return "", errors.New("key not found")
	}
	return "http://love_nika", nil /// как понимать длинная или нет я не ебу

}
func (s *MStorage) Ping() error {
	return errors.New("qwerty")
}
func (s *MStorage) AddM(mas []URLRegistryM, short []string, name string) error {
	return nil
}

func (s *MStorage) Find(oru string) (string, error) {
	return "", errors.New("")
}

func (s *MStorage) ListUserURLs(name string) []UserURL {
	return nil
}
