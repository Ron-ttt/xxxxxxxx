package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DBStorage struct {
	conn *pgx.Conn
}

type URL struct {
	ShortURL    string
	OriginalURL string
}

func NewDBStorage(dbname string) (Storage, error) {
	fmt.Println("3")
	conn, err := pgx.Connect(context.Background(), dbname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//defer conn.Close(context.Background())

	_, err1 := conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS hui(id SERIAL PRIMARY KEY, users text, shorturl text, originalurl text UNIQUE)")

	if err1 != nil {
		fmt.Println(err1)
		return nil, err1
	}
	fmt.Println("2")
	return &DBStorage{conn}, nil
}

func (s *DBStorage) Add(key string, value string, name string) error {
	_, err := s.conn.Exec(context.Background(), "INSERT INTO hui (shorturl, originalurl, users) VALUES($1, $2, $3)", key, value, name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *DBStorage) Get(key string) (string, error) {
	rows := s.conn.QueryRow(context.Background(), "SELECT originalurl FROM hui WHERE shorturl= $1", key)
	var originalURL string
	err := rows.Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

func (s *DBStorage) Ping() error {
	err := s.conn.Ping(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (s *DBStorage) AddM(m []URLRegistryM, short []string, name string) error {
	tx, err := s.conn.Begin(context.Background())
	if err != nil {
		return err
	}
	l := len(m)
	for i := 0; i < l; i++ {
		_, err := tx.Exec(context.Background(), "INSERT INTO hui (shorturl, originalurl, users)"+" VALUES($1,$2,$3)", short[i], m[i].OriginalURL, name)
		if err != nil {
			// если ошибка, то откатываем изменения
			tx.Rollback(context.Background())
			return err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (s *DBStorage) Find(oru string) (string, error) {
	rows := s.conn.QueryRow(context.Background(), "SELECT shorturl FROM hui WHERE originalurl= $1", oru)
	var short string
	err := rows.Scan(&short)
	if err != nil {
		return "", err
	}
	return short, nil
}

func (s *DBStorage) ListUserURLs(name string) []UserURL {
	var rez []UserURL

	rows := s.conn.QueryRow(context.Background(), "SELECT originalurl, shorturl FROM hui WHERE users= $1", name)
	err := rows.Scan(&rez)
	if err != nil {
		return nil
	}
	return rez
}
