package handlers

import (
	"EmployeeManagementSystem/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

	// Read the id parameter
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// Find the employee by id
	var employee models.Employee

	if result := h.DB.First(&employee, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Show the response

	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(employee)

	// for _, item := range mocks.Employees {
	// 	if item.Id == id {
	// 		w.Header().Set("ContentType", "application/json")
	// 		json.NewEncoder(w).Encode(item)
	// 		break
	// 	}
	// }
}
