package handlers

import (
	"EmployeeManagementSystem/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	// Read the id parameter
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	/*
	 - Read request body
	 - Defer statement is used to execute a function call just before the surrounding function where the defer statement is present returns.
	*/
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// Unmarshal parshes the json-encoded-data
	var update_employee models.Employee
	json.Unmarshal(body, &update_employee)

	// Find the employee by Id
	var employee models.Employee

	if result := h.DB.First(&employee, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Update the employee records
	employee.FullName = update_employee.FullName
	employee.Age = update_employee.Age

	h.DB.Save(&employee)

	// Show the response
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(employee)

}
