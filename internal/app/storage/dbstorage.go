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

	_, err1 := conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS hui(id SERIAL PRIMARY KEY,shorturl text, originalurl text UNIQUE)")
	fmt.Println("2")
	if err1 != nil {
		fmt.Println(err1)
		return nil, err1
	}
	return &DBStorage{conn}, nil
}

func (s *DBStorage) Add(key string, value string) error {
	_, err := s.conn.Exec(context.Background(), "INSERT INTO hui (shorturl, originalurl) VALUES($1, $2)", key, value)
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

func (s *DBStorage) AddM(m []URLRegistryM, short []string) error {
	tx, err := s.conn.Begin(context.Background())
	if err != nil {
		return err
	}
	l := len(m)
	for i := 0; i < l; i++ {
		_, err := tx.Exec(context.Background(), "INSERT INTO hui (shorturl, originalurl)"+" VALUES($1,$2)", short[i], m[i].OriginalURL)
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

func (s *DBStorage) GetU(name string) []UserURL {
	return nil
}
