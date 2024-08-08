package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBStorage struct {
	conn *pgxpool.Pool
}
type URL struct {
	ShortURL    string
	OriginalURL string
}

func NewDBStorage(dbname string) (Storage, error) {
	fmt.Println("3")
	conn, err := pgxpool.New(context.Background(), dbname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//defer conn.Close(context.Background())

	_, err1 := conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS urls(id SERIAL PRIMARY KEY, users text, shorturl text, originalurl text UNIQUE, isDeleted bool default false)")

	if err1 != nil {
		fmt.Println(err1)
		return nil, err1
	}
	fmt.Println("2")
	return &DBStorage{conn}, nil
}

func (s *DBStorage) Add(key string, value string, name string) error {
	_, err := s.conn.Exec(context.Background(), "INSERT INTO urls (shorturl, originalurl, users) VALUES($1, $2, $3)", key, value, name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *DBStorage) Get(key string) (string, error) {
	rows, err := s.conn.Query(context.Background(), "SELECT originalurl, isDeleted FROM urls WHERE shorturl= $1", key)
	if err != nil {
		return "", err
	}
	var originalURL string
	var del bool
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&originalURL, &del)
		if err != nil {
			return "", err
		}
	}
	if !del {
		return originalURL, nil
	} else {
		errurldeleted := errors.New("1")
		return "", errurldeleted
	}
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
		_, err := tx.Exec(context.Background(), "INSERT INTO urls (shorturl, originalurl, users)"+" VALUES($1,$2,$3)", short[i], m[i].OriginalURL, name)
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
	rows := s.conn.QueryRow(context.Background(), "SELECT shorturl FROM urls WHERE originalurl= $1", oru)
	var short string
	err := rows.Scan(&short)
	if err != nil {
		return "", err
	}
	return short, nil
}

func (s *DBStorage) ListUserURLs(name string) ([]UserURL, error) {
	var rez []UserURL
	rows, err := s.conn.Query(context.Background(), "SELECT originalurl, shorturl FROM urls WHERE users=$1", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rez1 UserURL
		err := rows.Scan(&rez1.OriginalURL, &rez1.ShortURL)
		if err != nil {
			return nil, err
		}
		rez = append(rez, rez1)
	}
	return rez, nil
}

func (s *DBStorage) DeleteURL(user string, short string) error {
	_, err1 := s.conn.Exec(context.Background(), "UPDATE urls SET isDeleted=TRUE WHERE shorturl=$1 AND users=$2", short, user)
	if err1 != nil {
		return err1
	}
	return nil
}
