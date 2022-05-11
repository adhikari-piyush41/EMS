package db

import (
	"log"

	"EmployeeManagementSystem/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize the database with gorm and db type
func Init() *gorm.DB {
	URL := "postgres://postgres:130310@localhost:5432/EmployeesRecord"

	// Opening the database connection
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})

	// Handling error
	if err != nil {
		log.Fatalln(err)
	}

	/* Migrate the schema to keep them up to date. AutoMigrate will create tables, missing foreign keys,
	constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed.
	It WON’T delete unused columns to protect your data.
	*/
	db.AutoMigrate(&models.Employee{})
	return db
}
