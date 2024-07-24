package storage

import (
	"encoding/json"
	"errors"
	"os"
)

type FileStorage struct {
	file          *os.File
	memoryStorage Storage
}

type FileJSON struct {
	UUID        int    `json:"uuid"`
	Users       string `json:"user"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewFileStorage(filename string) (Storage, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	memoryStorage := NewMapStorage()
	for decoder.More() {
		var data FileJSON
		decoder.Decode(&data)
		memoryStorage.Add(data.ShortURL, data.OriginalURL, data.Users)
	}

	return &FileStorage{
		file:          f,
		memoryStorage: memoryStorage,
	}, nil
}

func (s *FileStorage) Add(key string, value string, name string) error {
	s.memoryStorage.Add(key, value, name)
	data := FileJSON{UUID: 1, Users: name, ShortURL: key, OriginalURL: value}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(s.file.Name(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(append(jsonData, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStorage) Get(key string) (string, error) {
	return s.memoryStorage.Get(key)
}

func (s *FileStorage) Ping() error {
	return errors.New("тут нет бд")
}

func (s *FileStorage) AddM(mas []URLRegistryM, short []string, name string) error {
	l := len(mas)
	for i := 0; i < l; i++ {
		s.Add(mas[i].OriginalURL, short[i], name)
	}
	return nil
}

func (s *FileStorage) Find(oru string) (string, error) {
	return "", errors.New("")
}

func (s *FileStorage) ListUserURLs(name string) ([]UserURL, error) {
	return s.memoryStorage.ListUserURLs(name)
}
