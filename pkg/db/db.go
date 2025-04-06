package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

//дописать создание таблицы в бд
const (
	schema = `
	CREATE TABLE scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "", 
		title VARCHAR(32) NOT NULL DEFAULT "",
		comment TEXT NOT NULL DEFAULT "",
		repeat CHAR(8) NOT NULL DEFAULT ""
	);
	CREATE INDEX idx_data ON scheduler (date);`
)

func InitDB() error {
	dbFile := os.Getenv("TODO_DBFILE")

	var err error

	_, err = os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	var db *sqlx.DB

	if install {
		
		db, err = sqlx.Open("sqlite", dbFile)
		if err != nil {
			return fmt.Errorf("DB is not open: %v", err)
		}


		_, err = db.Exec(schema)
		if err != nil {
			return fmt.Errorf("Error create exec: %v", err)
		}

		log.Printf("A database with the scheduler table has been created, the path to the database: %v", dbFile)
	}

	db, err = sqlx.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("DB is not open: %v", err)
	}

	return nil
}