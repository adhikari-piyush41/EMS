// The module is created to initialize a single database connection instead of many.

package handlers

import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}
