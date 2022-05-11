package handlers

import (
	"EmployeeManagementSystem/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	// Read id parameter
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// Find the employee by Id
	var employee models.Employee

	if result := h.DB.First(&employee, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete the employee found
	h.DB.Delete(&employee)

	json.NewEncoder(w).Encode(employee)

	// for index, item := range mocks.Employees {
	// 	if item.Id == id {
	// 		// Deleting elements
	// 		mocks.Employees = append(mocks.Employees[:index], mocks.Employees[index+1:]...)
	// 		json.NewEncoder(w).Encode(item)
	// 		break
	// 	}
	// }
}
