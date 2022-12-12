package handlefuncs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/humberto1212/go-refreshment/models"
	"github.com/humberto1212/go-refreshment/psqlDb"
)

//===================
// overview of the API
//===================
// func Overview(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	http.ServeFile(w, r, "./static/index.html")
// }

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

	id := mux.Vars(r)["id"]

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

//============================
// 		Delete customer
//============================
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	db := psqlDb.Connect()
	defer db.Close()
	//check db
	errdb := db.Ping()
	if errdb != nil {
		panic(errdb)
	}

	rows, err := db.Query(`DELETE FROM customers where id=$1`, id)
	if err != nil {

		panic(err)
	}

	customer := models.Customer{}

	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Role, &customer.Email, &customer.Phone, &customer.Contacted)

		if err != nil {
			fmt.Println(err)
		}
	}

	w.WriteHeader(http.StatusOK)
}

//============================
// 		Create customer
//============================
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := psqlDb.Connect()
	defer db.Close()
	//check db
	errdb := db.Ping()
	if errdb != nil {
		panic(errdb)
	}

	var customer map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c_name := customer["name"].(string)
	c_role := customer["role"].(string)
	c_email := customer["email"].(string)
	c_phone := customer["phone"].(string)
	c_contactes := customer["contacted"].(bool)
	c_id := customer["id"].(float64)

	insertCustomer := `INSERT INTO customers(id, "name", "role", "email", "phone", "contacted") VALUES($1, $2, $3, $4, $5, $6)`
	_, insertError := db.Exec(insertCustomer, c_id, c_name, c_role, c_email, c_phone, c_contactes)
	if insertError != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

}

//============================
// 		Updatecs customer
//============================
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	db := psqlDb.Connect()
	defer db.Close()
	//check db
	errdb := db.Ping()
	if errdb != nil {
		panic(errdb)
	}

	var customer map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c_name := customer["name"].(string)
	c_role := customer["role"].(string)
	c_email := customer["email"].(string)
	c_phone := customer["phone"].(string)
	c_contacted := customer["contacted"].(bool)

	updateCustomer := `UPDATE customers SET name=$2, role=$3, email=$4, phone=$5, contacted=$6 WHERE id=$1;`
	_, updateError := db.Exec(updateCustomer, id, c_name, c_role, c_email, c_phone, c_contacted)
	if updateError != nil {
		panic(err)
	}

	errJson := json.NewEncoder(w).Encode(customer)
	if errJson != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
