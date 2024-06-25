package storage

import (
	"context"

	"github.com/jackc/pgx"
)

type DbStorage struct {
	conn *pgx.Conn
}

type URL struct {
	ShortURL    string
	OriginalURL string
}

func NewDbStorage(dbname string) (Storage, error) {
	conn, err := pgx.Connect(context.Background(), dbname)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return &DbStorage{conn}, nil
}

func (s *DbStorage) Add(key string, value string) error {
	_, err := s.conn.Exec("INSERT INTO $1 (shorturl, originalurl) VALUES($2, $3)", s.conn, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *DbStorage) Get(key string) (string, error) {
	rows := s.conn.QueryRow("SELECT short_url FROM $2 WHERE original_url = $1", s.conn, key)
	var originalURL string
	err := rows.Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

func (s *DbStorage) Ping() error {
	conn, err := pgx.Connect(context.Background(), s.conn)
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Ping(context.Background())
	if err != nil {
		return err
	}
	return nil
}
