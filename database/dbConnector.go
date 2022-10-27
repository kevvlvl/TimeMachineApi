package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DbUsername = "postgres"
	DbPassword = "timemachinedb"
	DbSchema   = "TimeMachine"
)

func OpenConnection() *sql.DB {

	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DbUsername, DbPassword, DbSchema)
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		fmt.Println("ERROR trying to connect to PostgresSQL using connection string. Error = ", err)
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
