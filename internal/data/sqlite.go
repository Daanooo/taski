package data

import (
	"database/sql"
	"errors"
	"io"
	"os"

	_ "modernc.org/sqlite"
)

func GetSqlite(name string) (*sql.DB, error) {
	shouldMigrate := false

	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		os.Create(name)
		shouldMigrate = true
	}

	db, err := sql.Open("sqlite", name)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	if shouldMigrate {
		err := migrate(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	// Not sure if the should be kept in the application directory, it could be altered by smart people that way, breaking the application
	// Maybe this can be done differently, I will look at that somewhere in the future
	file, err := os.Open("./dump.sql")
	if err != nil {
		return err
	}

	sql, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	db.Exec(string(sql))

	return nil
}
