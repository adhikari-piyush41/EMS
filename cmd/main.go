package main

import (
	"EmployeeManagementSystem/pkg/db"
	"EmployeeManagementSystem/pkg/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	r := mux.NewRouter()
	r.HandleFunc("/employees", h.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", h.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employees/add", h.AddEmployee).Methods("POST")
	r.HandleFunc("/employees/delete/{id}", h.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employees/update/{id}", h.UpdateEmployee).Methods("PUT")

	fmt.Printf("Starting server at port 4444\n")
	log.Fatal(http.ListenAndServe(":4444", r))
}
