package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"university-project/Model"
)

func getStatus(w http.ResponseWriter, r * http.Request) {
	// If Order Arrived Successfully will show this response
	Arrived := Model.Status{Code:201,Message:"Your Order Arrived"}
	// If Order Didn't Arrive will show this response
	NotArrived := Model.Status{Code:401,Message:"Your Order in Progress"}

	id := r.Header.Get("id")
	// Database implementation
	db, err := gorm.Open("sqlite3", "university.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Operation Model
	var operations Model.Operation
	db.Where("ID = ? ", id).Find(&operations)
	if operations.Arrived != "true" {
		json.NewEncoder(w).Encode(NotArrived)
	} else {
		json.NewEncoder(w).Encode(Arrived)
	}
}
