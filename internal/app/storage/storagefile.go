package storage

import (
	"encoding/json"
	"os"
)

type FileStorage struct {
	file          *os.File
	memoryStorage Storage
}

type FileJSON struct {
	UUID        int    `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewFileStorage(filename string) (Storage, error) {
	f, err := os.OpenFile("tmp/"+filename, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	memoryStorage := NewMapStorage()
	for decoder.More() {
		var data FileJSON
		decoder.Decode(&data)
		memoryStorage.Add(data.ShortURL, data.OriginalURL)
	}
	return &FileStorage{
		file:          f,
		memoryStorage: memoryStorage,
	}, nil
}

func (s *FileStorage) Add(key string, value string) error {
	s.memoryStorage.Add(key, value)
	data := FileJSON{UUID: 1, ShortURL: key, OriginalURL: value}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(s.file.Name(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	_, err = f.Write(append(jsonData, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStorage) Get(key string) (string, error) {
	return s.memoryStorage.Get(key)
}
