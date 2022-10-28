package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DbHost     = "localhost"
	DbPort     = 5432
	DbUsername = "postgres"
	DbPassword = "timemachinedb"
	DbName     = "postgres"
)

func OpenConnection() *sql.DB {

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUsername, DbPassword, DbName)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println("ERROR trying to connect to PostgresSQL using connection string. Error = ", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR trying to ping the database. Error = ", err)
		return nil
	}

	return db
}

func CloseConnection(dbInstance *sql.DB) {

	if dbInstance != nil {

		err := dbInstance.Close()

		if err != nil {

			fmt.Println("ERROR trying to close to PostgresSQL connection. Error = ", err)
			return
		}
	}
}
