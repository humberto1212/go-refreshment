package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/humberto1212/go-refreshment/handlefuncs"
	"github.com/humberto1212/go-refreshment/psqlDb"
)

func main() {

	//=====================
	//Connect to Data base
	//=====================
	db := psqlDb.Connect()
	defer db.Close()
	//check db
	errdb := db.Ping()
	if errdb != nil {
		panic(errdb)
	}

	//create table
	createTable := `CREATE TABLE IF NOT EXISTS Customers (ID INT PRIMARY KEY, Name TEXT, Role TEXT, Email TEXT, Phone TEXT, Contacted BOOL);`
	_, e := db.Exec(createTable)
	if e != nil {
		panic(e)
	}

	// ===============================
	// insert and delete
	// hardcoded values in the table
	// ===============================

	//Delete all rows
	deleteAllRowsStmt := `DELETE FROM customers;`
	_, e = db.Exec(deleteAllRowsStmt)
	if e != nil {
		panic(e)
	}

	//insert Rows
	insert1Stmt := `INSERT INTO "customers"(id, "name", "role", "email", "phone", "contacted") values(1, 'Jacob', 'Manager', 'jm@gmail.com', '0157-22244', true )`
	_, e = db.Exec(insert1Stmt)
	if e != nil {
		panic(e)
	}

	insert2Stmt := `INSERT INTO "customers"(id, "name", "role", "email", "phone", "contacted") values(2, 'Pedro', 'Marketing', 'P@gmail.com', '0157-13544', false )`
	_, e = db.Exec(insert2Stmt)
	if e != nil {
		panic(e)
	}

	insert3Stmt := `INSERT INTO "customers"(id, "name", "role", "email", "phone", "contacted") values(3, 'Fiore', 'Developer', 'F@gmail.com', '0157-22235', true )`
	_, e = db.Exec(insert3Stmt)
	if e != nil {
		panic(e)
	}

	//======================================
	//				Handlers
	//======================================
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers", handlefuncs.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handlefuncs.GetSingleCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", handlefuncs.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers", handlefuncs.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handlefuncs.UpdateCustomer).Methods("PUT")

	muxWithMiddlewares := http.TimeoutHandler(router, time.Second*5, "Timeout!")

	fmt.Println("Server is starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", muxWithMiddlewares))
}
