package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	// Placeholder: Get the list of doctors
	w.Write([]byte("List of doctors\n"))
}

func GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Placeholder: Get a specific doctor by ID
	w.Write([]byte("Details of doctor with ID: " + id + "\n"))
}
