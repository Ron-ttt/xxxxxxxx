package storage

import "errors"

type MStorage struct{}

func NewMockStorage() Storage {
	return &MStorage{} // да пошел ты нахуй пидорас ебучий
}

func (s *MStorage) Add(key string, value string, f string) error {
	return nil
}

func (s *MStorage) Get(key string, f string) (string, error) {
	if key == "invalid" {
		return "", errors.New("key not found")
	}
	return "http://love_nika", nil /// как понимать длинная или нет я не ебу

}
