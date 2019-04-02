package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io"
	"net/http"
	"os"
	"time"
	"university-project/Model"
)


// New Operation
func createOperation(w http.ResponseWriter, r *http.Request) {
	// Database implementation
	db, err := gorm.Open("sqlite3", "university.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()


	// Operation Model
	var operation Model.Operation


	// Check if MultipartForm not null will take data from it and ignore Body Data
	r.ParseMultipartForm(24 * 80 * 1080)

	// Create File
	file,handler, err := r.FormFile("File")
	defer file.Close()
	f, err := os.OpenFile("./files/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)


	/*
	** Note: FormFiles aren't added to FormFiles so we need to add it manually by this line
	*/
	r.Form.Add("File",handler.Filename)
	r.Form.Add("Arrived", "false")

	// Schema Encoder
	decoder := schema.NewDecoder()
	decoder.Decode(&operation, r.Form)


	// Save operation
	db.Create(&operation)


	// Show operation in response
	json.NewEncoder(w).Encode(&operation)

	time.AfterFunc(operation.Date * time.Second, func() {
		db, err := gorm.Open("sqlite3", "university.db")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()
		db.Model(&operation).Where("ID = ?", operation.ID).Update("Arrived", "true")
	})
}

// Get all operations
func getOperations(w http.ResponseWriter,r * http.Request) {
	db, err := gorm.Open("sqlite3", "university.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var operations []Model.Operation
	db.Find(&operations)
	json.NewEncoder(w).Encode(operations)
}

