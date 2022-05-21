// Package storage provides access to the database.
package storage

import (
	"database/sql"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

// Sqlite3 - лучший файлообменник.
const driver = "sqlite3_custom"

// Storage represents a storage.
type Storage struct {
	Voices *VoicesRepository
	Pastas *PastasRepository
}

// Repository represents a data repository.
type Repository interface {
	Init() error
}

// New creates a new storage.
func New(dsn string) (*Storage, error) {
	sql.Register("sqlite3_custom", &sqlite3.SQLiteDriver{
		ConnectHook: func(sc *sqlite3.SQLiteConn) error {
			if err := sc.RegisterFunc("similarity", func(a, b string) float64 {
				return strutil.Similarity(a, b, metrics.NewSmithWatermanGotoh())
			}, true); err != nil {
				return err
			}

			// Native lower does not work with Unicode characters.
			if err := sc.RegisterFunc("lower", func(s string) string {
				return strings.ToLower(s)
			}, true); err != nil {
				return err
			}

			return nil
		},
	})

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	s := &Storage{
		Voices: &VoicesRepository{db: db},
		Pastas: &PastasRepository{db: db},
	}
	for _, repo := range []Repository{s.Voices, s.Pastas} {
		if err = repo.Init(); err != nil {
			return nil, err
		}
	}

	return s, nil
}
