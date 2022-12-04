package psqlDb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	env_path = "/Users/humberto/go/src/github.com/humberto1212/ refreshment"
)

func Connect() *sql.DB {
	// godotenv package
	host := goDotEnvVariable("HOST")
	port := goDotEnvVariable("PORT")
	user := goDotEnvVariable("USER")
	dbname := goDotEnvVariable("DBNAME")

	// connection string
	psqlconst := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconst)
	CheckError(err)

	fmt.Println("Connected db")

	return db

}

//===========================
// 	Error handling
//===========================
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//===================================================
// use godot package to load/read the .env file and
//===================================================
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(filepath.Join(env_path, ".env"))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
