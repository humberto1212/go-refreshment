package handlefuncs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/humberto1212/go-refreshment/models"
	"github.com/humberto1212/go-refreshment/psqlDb"
)

func GetAllCustomers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("conten-type", "aplication/json")

	db := psqlDb.Connect()

	defer db.Close()

	// check db
	// errdb := db.Ping()
	// if errdb != nil {
	// 	panic(errdb)
	// }

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

	fmt.Println(customers)
	errJson := json.NewEncoder(w).Encode(customers)
	if errJson != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
