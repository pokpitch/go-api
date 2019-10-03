package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func getAllAddressBook(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook {
		Firstname:	"Pokpitch",
		Lastname:	"Patcharadamrongkul",
		Code:		1983,
		Phone:		"0898107096",
		}
		json.NewEncoder(w).Encode(addBook)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/address", getAllAddressBook)
	http.ListenAndServe(":8080", nil)
}

type addressBook struct {
    Firstname string
    Lastname  string
    Code      int
    Phone     string
}