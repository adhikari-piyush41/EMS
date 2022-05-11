package handlers

import (
	"EmployeeManagementSystem/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	// Find employee in database
	var employee []models.Employee

	if result := h.DB.Find(&employee); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Show response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
