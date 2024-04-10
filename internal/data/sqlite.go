package data

import (
	"database/sql"
	"errors"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetSqlite(name string) (*sql.DB, error) {
	shouldMigrate := false

	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		log.Printf("Created database %s\n", name)
		os.Create(name)
		shouldMigrate = true
	}

	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	if shouldMigrate {
		err := migrate(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	log.Println("Migrating database using dump.sql")

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
