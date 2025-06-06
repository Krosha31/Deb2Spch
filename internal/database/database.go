package database

import (
	"database/sql"
	"fmt"
	"os"

	. "Deb2Spch/internal/common"
	_ "github.com/lib/pq"
)


const (
    host     = "db" //localhost
    port     = 5432
    user     = "user"
    password = "1111"
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
	fmt.Println("executed")
	cwd, _ := os.Getwd()
	files, err := os.ReadDir(cwd + "/" + filePath)
	if err != nil {
		return err
	}
	for _, file := range files {
		sqlBytes, err := os.ReadFile(filePath + file.Name())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		_, err = database.db.Exec(string(sqlBytes))
		if err != nil {
			fmt.Println(err.Error())
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
	// database.connectToServer()
	// database.executeFunctions()
	// database.createDB()
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

	error := database.executeFunctions()
	fmt.Println(error)
	error = database.createAllTables()
	fmt.Println(error)
	database.AddSubscription(0)
	return nil
}


func (database *Database) Disconnect() {
	fmt.Println("Disconnected")
	database.db.Close()
}

func (database *Database) AddUser(login string, password string) error {
	query := "SELECT add_user($1, $2)"
	_, err := database.db.Exec(query, login, password)
    if err != nil {
        return err
    }
	return nil;
}

func (database *Database) AddSubscription(duration int) error {
	query := "SELECT add_suncription($1)"
	_, err := database.db.Exec(query, duration)
    if err != nil {
        return err
    }
	return nil;
}

func (database *Database) GetUserByLogin(login string) (User, error) {
	query := "SELECT * FROM get_user_by_login($1)"
	var user User
	err := database.db.QueryRow(query, login).Scan(&user.Email, &user.Password_hash, &user.Subscribtion_id, &user.Registration_date)
    if err != nil {
        if err == sql.ErrNoRows {
            return User{}, nil
        }
        return User{}, err
    }
	return user, nil;
}

func (database *Database) InsertRequest(userID, input string) (int, error) {
    var newID int
    query := "SELECT insert_request($1, $2)"
    err := database.db.QueryRow(query, userID, input).Scan(&newID)
    if err != nil {
        return 0, err
    }
    return newID, nil
}

func (database *Database) GetRequestsByUser(userID string) ([]Request, error) {
    query := "SELECT * FROM get_requests_by_user($1)"
    rows, err := database.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var requests []Request
    for rows.Next() {
        var req Request
        if err := rows.Scan(&req.ID, &req.UserID, &req.Time, &req.Input); err != nil {
            return nil, err
        }
        requests = append(requests, req)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return requests, nil
}


var Db Database;