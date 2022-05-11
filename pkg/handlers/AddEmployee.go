package handlers

import (
	"EmployeeManagementSystem/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Inside the function to access the database defining the function on the handler struct
func (h handler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var employee models.Employee
	json.Unmarshal(body, &employee)

	// Give a new id and strconv formats it to the string
	// employee.Id = rand.Intn(1000)
	// mocks.Employees = append(mocks.Employees, employee)

	// Add data
	if result := h.DB.Create(&employee); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send response
	w.Header().Set("ContentType", "application/json")

	json.NewEncoder(w).Encode(employee)
}
