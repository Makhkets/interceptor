package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// set CGO_ENABLED=1
// go env -w CGO_ENABLED=1

type storage struct {
	db *sql.DB
}

type Storage interface {
	InsertInformation(cookie, remoteAddr string) error
}

func New(storagePath string) (Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &storage{db: db}, nil
}

type StorageError error

var (
	UserNotFoundErr StorageError = errors.New("user not found")
)

func (s *storage) InsertInformation(cookie, remoteAddr string) error {
	const op = "storage.sqlite.InsertInformation"

	_, err := s.db.Exec(`INSERT INTO logs (cookie, ip) VALUES (?, ?)`, cookie, remoteAddr)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
