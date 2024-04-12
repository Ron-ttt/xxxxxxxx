package storage

import (
	"encoding/json"
	"errors"
	"os"
)

type Storage interface {
	Add(key string, value string, f string) error
	//Remove(key string)
	Get(key string, f string) (string, error)
}

type MapStorage struct {
	m map[string]string
}

var i int

type FileJ struct {
	Uuid        int    `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewMapStorage() Storage {
	return &MapStorage{
		m: make(map[string]string)}
}

func (s *MapStorage) Add(key string, value string, f string) error {
	if len(f) > 0 {
		addfile(key, value, f)
		s.m[key] = value
	}
	s.m[key] = value
	return nil
}

func addfile(key string, value string, f string) error {
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE, 0)
	if err != nil {
		return err
	}
	defer file.Close()
	var data = FileJ{Uuid: i, ShortURL: key, OriginalURL: value}
	d, _ := json.Marshal(data)
	d = append(d, '\n')
	_, err = file.Write(d)
	if err != nil {
		return err
	}
	i++
	return nil
}

// func (s *MapStorage) Remove(key string) {
// 	delete(s.m, key)
// }

func (s *MapStorage) Get(key string, f string) (string, error) {
	value, found := s.m[key]
	if !found {
		return "", errors.New("key not found")
	}
	return value, nil
}
