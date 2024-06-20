package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

type DbStorage struct {
}

type URL struct {
	ShortURL    string
	OriginalURL string
}

func NewDbStorage(dbname string) (Storage, error) {
	conn, err := pgx.Connect(context.Background(), dbAdress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	return &DbStorage{}, nil
}

func (s *DbStorage) Add(key string, value string) error {

	return nil
}

func (s *DbStorage) Get(key string) (string, error) {
	db, err := sql.Open("pgx", "hui")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT shorturl FROM hui WHERE originalurl = $1", key)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var v URL
	err = rows.Scan(&v.ShortURL)
	if err != nil {
		return "", err
	}

	return v.ShortURL, nil
}
