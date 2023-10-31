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

func (dController *PacientController) GetPacienti(w http.ResponseWriter, r *http.Request) {
	log.Println("[PACIENTI] Fetching all pacienti...")

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctors, err := dController.DbConn.FetchPacienti(ctx)
	if err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(doctors); err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Println("[PACIENTI] Successfully fetched all doctors")
}

func (dController *PacientController) GetPacientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid pacient ID: %s", idStr)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
		return
	}

	log.Printf("Fetching pacient with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	pacient, err := dController.DbConn.FetchPacientByID(ctx, id)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to fetch pacient with ID %d: %s", id, err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if pacient == nil {
		err_msg := fmt.Sprintf("No pacient found with ID: %d", id)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pacient); err != nil {
		err_msg := fmt.Sprintf("Error encoding pacient to JSON: %s", err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PACIENTI] Successfully fetched pacient with ID: %d", id)
}

func (dController *PacientController) GetPacientByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	log.Printf("[PACIENTI] Fetching pacient with email: %s...", email)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	pacient, err := dController.DbConn.FetchPacientByEmail(ctx, email)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to fetch pacient with email %s: %s", email, err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if pacient == nil {
		err_msg := fmt.Sprintf("No pacient found with email: %s", email)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pacient); err != nil {
		err_msg := fmt.Sprintf("Error encoding pacient to JSON: %s", err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PACIENTI] Successfully fetched pacient with email: %s", email)
}

func (dController *PacientController) GetPacientByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars["id"]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
		return
	}

	log.Printf("[PACIENTI] Fetching pacient with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	pacient, err := dController.DbConn.FetchPacientByUserID(ctx, userID)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to fetch pacient with user ID %d: %s", userID, err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if pacient == nil {
		err_msg := fmt.Sprintf("No pacient found with user ID: %d", userID)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pacient); err != nil {
		err_msg := fmt.Sprintf("Error encoding pacient to JSON: %s", err)
		log.Printf("[PACIENTI] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PACIENTI] Successfully fetched pacient with user ID: %d", userID)
}
