package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	db *sql.DB
}

func NewDatabaseConnection(username, password string) *DatabaseConnection {

	connStr := fmt.Sprintf("host=localhost user=%v password=%v dbname=helpdesk sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &DatabaseConnection{
		db: db,
	}

}

func (dc DatabaseConnection) Init() (*sql.DB, error) {

	err := dc.createTables()
	if err != nil {
		return nil, err
	}
	return dc.db, nil
}

func (dc DatabaseConnection) createTables() error {

	_, err := dc.db.Exec(userTable)
	if err != nil {
		return err
	}
	_, err = dc.db.Exec(ticketTable)
	if err != nil {
		return err
	}
	return nil
}
