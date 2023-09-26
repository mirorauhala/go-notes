package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {

	// read json file
	jsonFile, err := os.Open("database.json")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error reading json file"}`))
		return
	}

	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue := make([]byte, 1000)
	jsonFile.Read(byteValue)

	// unmarshal byte array to struct
	type User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	type Result struct {
		Users []User `json:"users"`
	}

	var result Result
	errr := json.Unmarshal(byteValue, &result)
	if errr != nil {
		// Handle the error
		fmt.Println("Error:", errr)
		return
	}

	w.Write([]byte("foo"))
}
