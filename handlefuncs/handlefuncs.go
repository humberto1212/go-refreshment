package handlefuncs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/humberto1212/go-refreshment/models"
	"github.com/humberto1212/go-refreshment/psqlDb"
)

//===================
// Get all Customers
//===================
func GetAllCustomers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("conten-type", "aplication/json")

	db := psqlDb.Connect()

	defer db.Close()

	//check db
	errdb := db.Ping()
	if errdb != nil {
		panic(errdb)
	}

	var customers []models.Customer

	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		customer := models.Customer{}

		err = rows.Scan(&customer.ID, &customer.Name, &customer.Role, &customer.Email, &customer.Phone, &customer.Contacted)
		if err != nil {
			fmt.Println(err)
		}

		customers = append(customers, customer)
	}

	errJson := json.NewEncoder(w).Encode(customers)
	if errJson != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

//============================
// Getting a single customer
//============================
func GetSingleCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("conten-type", "aplication/json")

	idStirng := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStirng)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	fmt.Println("++++++++++++++++///+>", id)

	db := psqlDb.Connect()
	defer db.Close()
	//check db
	errdb := db.Ping()
	if errdb != nil {
		panic(errdb)
	}

	customer := models.Customer{}

	rows, err := db.Query(`SELECT id, name, role, email, phone, contacted  FROM customers where id=$1`, id)
	if err != nil {

		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&customer.ID, &customer.Name, &customer.Role, &customer.Email, &customer.Phone, &customer.Contacted)

		if err != nil {
			fmt.Println(err)
		}

	}

	errJson := json.NewEncoder(w).Encode(customer)
	if errJson != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
