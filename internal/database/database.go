package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)


const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "0793"
    dbname   = "Deb2Spch"
	filePath = "migrations/sql/"
)

type Database struct {
	db *sql.DB
}

func (database *Database) NewDatabase() {
	;
}

func (database *Database) executeFunctions() error {
	files, err := os.ReadDir(filePath)
	if err != nil {
		return err
	}
	for _, file := range files {
		sqlBytes, err := os.ReadFile(filePath + file.Name())
		if err != nil {
			return err
		}
		_, err = database.db.Exec(string(sqlBytes))
		if err != nil {
			return err
		}
	}
	return nil
}

func (database *Database) createDB() error {
	query := "SELECT create_db($1)"
	_, err := database.db.Exec(query, dbname)
    if err != nil {
        return err
    }
	fmt.Printf("Database '%s' created successfully\n", dbname)
	return nil
}

func (database *Database) createAllTables() error {
	query := "SELECT create_all_tables()"
	_, err := database.db.Exec(query)
    if err != nil {
        return err
    }
	fmt.Println("Tables created successfully")
	return err
}


func (database *Database) connectToServer() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
	host, port, user, password)
	var err error
	database.db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}
	err = database.db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to the PostgreSQL server")
	return nil
}

func (database *Database) Connect() error {
	database.connectToServer()
	database.executeFunctions()
	database.createDB()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	var err error
	database.db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}
	err = database.db.Ping()
	if err != nil {
		return err
	}

	fmt.Printf("Connected to the %s\n", dbname)

	database.executeFunctions()
	err = database.createAllTables()
	fmt.Println(err)
	return nil
}


func (database *Database) Disconnect() {
	fmt.Println("Disconnected")
	database.db.Close()
}