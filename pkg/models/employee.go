package models

// Creating a struct for storing multiple values of different data types.
type Employee struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}
