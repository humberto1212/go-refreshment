package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/humberto1212/go-refreshment/handlefuncs"
)

func main() {

	// // check db
	// err = db.Ping()
	// CheckError(err)

	// create table
	// createTable := `CREATE TABLE IF NOT EXISTS Customers (ID INT PRIMARY KEY, Name TEXT, Role TEXT, Email TEXT, Phone TEXT, Contacted BOOL);`
	// _, e := db.Exec(createTable)
	// CheckError(e)

	// ===============================
	// insert
	// hardcoded values in the table
	// ===============================
	// insertStmt := `insert into "customers"(id, "name", "role", "email", "phone", "contacted") values(1, 'Jacob', 'Manager', 'jm@gmail.com', '0157-22244', true )`
	// _, e = db.Exec(insertStmt)
	// CheckError(e)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/allCustomers", handlefuncs.GetAllCustomers).Methods("GET")
	router.HandleFunc("/allCustomers/{id}", handlefuncs.GetSingleCustomer).Methods("GET")

	fmt.Println("Server is starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
