package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// Retrieve all programari
func (pController *ProgramareController) GetProgramari(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PROGRAMARE] Attempting to retrieve programari.")

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to fetch all programari from the database
	programari, err := pController.DbConn.FetchProgramari(ctx, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to fetch programari: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PROGRAMARE] Successfully fetched all programari")
	// Serialize the programari to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, programari)
}

// Retrieve a programare by ID
func (pController *ProgramareController) GetProgramareByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PROGRAMARE] Attempting to retrieve a programare by ID.")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid programare ID", http.StatusBadRequest)
		return
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to fetch the programare by ID from the database
	programare, err := pController.DbConn.FetchProgramareByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to fetch programare by ID: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// Check if the programare exists
	if programare == nil {
		http.Error(w, "Programare not found", http.StatusNotFound)
		return
	}

	log.Printf("[PROGRAMARE] Successfully fetched programare %d", programare.IDProgramare)
	// Serialize the programare to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, programare)
}

// Retrieve programari by doctor ID
func (pController *ProgramareController) GetProgramariByDoctorID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PROGRAMARE] Attempting to retrieve programari by Doctor ID.")
	vars := mux.Vars(r)
	doctorID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Doctor ID", http.StatusBadRequest)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to fetch programari by Doctor ID from the database
	programari, err := pController.DbConn.FetchProgramariByDoctorID(ctx, doctorID, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to fetch programari by Doctor ID: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PROGRAMARE] Successfully fetched programari of doctor %d", doctorID)
	// Serialize the programari to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, programari)
}

// Retrieve programari by pacient ID
func (pController *ProgramareController) GetProgramariByPacientID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PROGRAMARE] Attempting to retrieve programari by Pacient ID.")
	vars := mux.Vars(r)
	pacientID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Pacient ID", http.StatusBadRequest)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to fetch programari by Pacient ID from the database
	programari, err := pController.DbConn.FetchProgramariByPacientID(ctx, pacientID, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to fetch programari by Pacient ID: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PROGRAMARE] Successfully fetched programari of pacient %d", pacientID)
	// Serialize the programari to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, programari)
}

// Retrieve programari by date
func (pController *ProgramareController) GetProgramariByDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PROGRAMARE] Attempting to retrieve programari by date.")

	// Use mux.Vars to get the date parameter from the route
	vars := mux.Vars(r)
	dateStr := vars["date"]

	// Parse the date string into a time.Time object
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Printf("[PROGRAMARE] Failed to convert date string: %s\n", err)
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to fetch programari by date from the database
	programari, err := pController.DbConn.FetchProgramariByDate(ctx, date, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to fetch programari by date: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PROGRAMARE] Successfully fetched programari from %s", date)
	// Serialize the programari to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, programari)
}

// Retrieve programari by status
func (pController *ProgramareController) GetProgramariByStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PROGRAMARE] Attempting to retrieve programari by status.")

	// Use mux.Vars to get the status parameter from the route
	vars := mux.Vars(r)
	status := vars["status"]

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to fetch programari by status from the database
	programari, err := pController.DbConn.FetchProgramariByStatus(ctx, status, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to fetch programari by status: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PROGRAMARE] Successfully fetched programari with status of %s", status)
	// Serialize the programari to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, programari)
}
