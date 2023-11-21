package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	db *sql.DB
}

func DBConnect(fileName string) *sqliteStore {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to connect to DB: %s", fileName))
	}

	createMainTables := `
	create table if not exists category (
		id INTEGER PRIMARY KEY, 
		name varchar(50) not null 
	);
	` + GetTransactionTable()

	_, err = db.Exec(createMainTables)
	if err != nil {
		log.Fatalln("Could not create main tables, error:", err)
	}

	return &sqliteStore{db: db}
}

func (s *sqliteStore) Close() {
	s.db.Close()
}

func (s *sqliteStore) GetCategoryID(name string) (int, error) {
	query := "SELECT id FROM category WHERE name = ?"
	var categoryID int
	err := s.db.QueryRow(query, name).Scan(&categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, err
		}
		log.Fatalln("Failed to query category table, error:", err)
	}

	return categoryID, nil
}

func (s *sqliteStore) CreateCategory(name string) int {
	query := "INSERT INTO category (name) VALUES (?)"
	result, err := s.db.Exec(query, name)
	if err != nil {
		log.Fatalln("Could not insert into category table, error:", err)
	}

	categoryID, err := result.LastInsertId()
	if err != nil {
		log.Fatalln("Could not retrive id from generated category field, error:")
	}

	return int(categoryID)
}

func (s *sqliteStore) AddTransaction(transaciton Transaction) error {
	query := "INSERT INTO transaction (catergory_id, amount, title) VALUES (?, ?, ?)"
	_, err := s.db.Exec(query, transaciton.CategoryID, transaciton.Amount, transaciton.Title)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteStore) PrintCategory() {
	rows, err := s.db.Query("select * from category")
	if err != nil {
		log.Fatalln("failed to select *, error:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id any
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatalln("failed to scan, error:", err)
		}
		fmt.Println("Category:", id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln("rows errored, error:", err)
	}
}
