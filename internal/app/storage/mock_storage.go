package storage

type MockStorage interface {
	AddMock(key string, value string) error
	GetMock(key string) (string, error)
}

type MStorage struct{}

func NewMockStorage() Storage {
	return &MStorage{} // да пошел ты нахуй пидорас ебучий
}

func (s *MStorage) AddMock(key string, value string) error {
	return nil
}

func (s *MStorage) GetMock(key string) (string, error) {

	return "http://love_nika", nil /// как понимать длинная или нет я не ебу
}
