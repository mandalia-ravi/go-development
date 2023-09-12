package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS employees (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/employees", getEmployees(db)).Methods("GET")
	router.HandleFunc("/employees/{id}", getEmployee(db)).Methods("GET")
	router.HandleFunc("/employees", createEmployee(db)).Methods("POST")
	router.HandleFunc("/employees/{id}", updateEmployee(db)).Methods("PUT")
	router.HandleFunc("/employees/{id}".deleteEmployee(db)).Methods("DELETE")

}

// get all employees
func getEmployees(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// get employee by id
func getEmployee(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func createEmployee(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e Employee
		json.NewDecoder(r.Body).Decode(&e)

		err := db.QueryRow("INSERT INTO employees (name, email) VALUES ($1, $2) RETURNING id", e.Name, e.Email).Scan(&e.ID)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(e)
	}
}

func updateEmployee(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func deleteEmployee(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
