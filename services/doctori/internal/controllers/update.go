package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Placeholder: Update doctor with given ID
	w.Write([]byte("Updated doctor with ID: " + id + "\n"))
}
