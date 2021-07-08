package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err := sql.Open("sqlite3", "./studybuddy.db")

	if err != nil {
		fmt.Printf("errore %t", err)
		log.Fatal(err.Error())
		return err
	}
	//defer db.Close()

	return db.Ping()
}

func InsertNote(word string, definition string, category string) error {
	db, err := sql.Open("sqlite3", "./studybuddy.db")

	if err != nil {
		fmt.Printf("errore %t", err)
		log.Fatal(err.Error())
		return err
	}
	InsertNoteSQL := `INSERT INTO studybuddy (word, definition, category) VALUES (?, ?, ?)`

	statement, err := db.Prepare(InsertNoteSQL)

	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	log.Println("Nota inserita correttamente")
	return nil
}

func CreateTable() {
	fmt.Println("CreateTable STARTED")
	db, err := sql.Open("sqlite3", "./studybuddy.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS studybuddy (
        "idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "word" TEXT,
        "definition" TEXT,
        "category" TEXT
      );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Studybuddy table created")

}
