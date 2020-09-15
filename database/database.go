package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	// this is required to enable database interaction. It calls a construction function
	_ "github.com/lib/pq"
)

// Init .. initialises the database tables
func Init(sqlStatement string) {

	// create the postgres db connection
	db := Connect()
	defer db.Close()

	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
}

// Connect gets connection details from env variables and returns a pointer to the database
func Connect() (db *sql.DB) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		port = 5432
		log.Println("Error setting PORT from env, default used.")
	}

	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "postgresql.server"
		log.Println("Error setting HOST from env, default used.")
	}
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "postgres"
		log.Println("Error setting USER from env, default used.")
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {

		password = "password"
		log.Println("Error setting PASSWORD from env, default used.")
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		dbname = "connolly_bank"
		log.Println("Error setting DATABASE from env, default used.")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	fmt.Println(psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to Database %s", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging Database %s", err)
		return
	}

	fmt.Println("Postgres connected successfully")
	return
}
