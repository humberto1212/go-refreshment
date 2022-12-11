# Customer relationship management Backend with Golang

The project represents the backend of a customer relationship management (CRM) web application.

## Specifications

### The project was created with:
1. Golang

2. DB:
- PSQL

### Project features

- Getting a list of all customers
- Getting data for a single customer
- Adding a customer
- Updating a customer's information
- Removing a customer

### Instructions:
- [install PostgreSQL ](https://www.postgresql.org/download/)

- In the terminal run `go run main.go`, this start the server.

- Open (http://localhost:8080) or use [postman](https://www.postman.com/)

- how to interact with the server?
Getting a single customer through a `/customers/{id}` path
Getting all customers through a the `/customers` path
Creating a customer through a `/customers` path
Updating a customer through a `/customers/{id}` path
Deleting a customer through a `/customers/{id}` path

-  create a new customer 
`
 {
     "id": 4,
     "name": "Carola",
     "role": "Manager",
     "email": "C@gmail.com",
     "phone": "12345",
     "contacted": false
}
`