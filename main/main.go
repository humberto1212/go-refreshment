package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
)

func main() {
	// connection string
	psqlconst := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconst)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
