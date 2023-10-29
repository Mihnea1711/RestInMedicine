package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (dController *DoctorController) GetDoctors(w http.ResponseWriter, r *http.Request) {
	doctors, err := dController.DbConn.FetchDoctors()
	if err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(doctors); err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}
}

func GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Placeholder: Get a specific doctor by ID
	w.Write([]byte("Details of doctor with ID: " + id + "\n"))
}
