package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (dController *DoctorController) GetDoctors(w http.ResponseWriter, r *http.Request) {
	log.Println("[DOCTOR] Fetching all doctors...")

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctors, err := dController.DbConn.FetchDoctors(ctx)
	if err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(doctors); err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Println("[DOCTOR] Successfully fetched all doctors")
}

func (dController *DoctorController) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid doctor ID: %s", idStr)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByID(ctx, id)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to fetch doctor with ID %d: %s", id, err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if doctor == nil {
		err_msg := fmt.Sprintf("No doctor found with ID: %d", id)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(doctor); err != nil {
		err_msg := fmt.Sprintf("Error encoding doctor to JSON: %s", err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Printf("[DOCTOR] Successfully fetched doctor with ID: %d", id)
}

func (dController *DoctorController) GetDoctorByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	log.Printf("[DOCTOR] Fetching doctor with email: %s...", email)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByEmail(ctx, email)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to fetch doctor with email %s: %s", email, err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if doctor == nil {
		err_msg := fmt.Sprintf("No doctor found with email: %s", email)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(doctor); err != nil {
		err_msg := fmt.Sprintf("Error encoding doctor to JSON: %s", err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Printf("[DOCTOR] Successfully fetched doctor with email: %s", email)
}

func (dController *DoctorController) GetDoctorByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars["id"]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByID(ctx, userID)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to fetch doctor with user ID %d: %s", userID, err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if doctor == nil {
		err_msg := fmt.Sprintf("No doctor found with user ID: %d", userID)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(doctor); err != nil {
		err_msg := fmt.Sprintf("Error encoding doctor to JSON: %s", err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Printf("[DOCTOR] Successfully fetched doctor with user ID: %d", userID)
}
