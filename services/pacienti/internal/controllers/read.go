package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (dController *PacientController) GetPacienti(w http.ResponseWriter, r *http.Request) {
	log.Println("[PATIENT] Fetching all pacienti...")

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	pacients, err := dController.DbConn.FetchPacienti(ctx)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, pacients)
}

func (dController *PacientController) GetPacientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid pacient ID: %s", idStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	log.Printf("Fetching pacient with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	pacient, err := dController.DbConn.FetchPacientByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch pacient with ID %d: %s", id, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if pacient == nil {
		errMsg := fmt.Sprintf("No pacient found with ID: %d", id)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, pacient)
}

func (dController *PacientController) GetPacientByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	log.Printf("[PATIENT] Fetching pacient with email: %s...", email)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	pacient, err := dController.DbConn.FetchPacientByEmail(ctx, email)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch pacient with email %s: %s", email, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if pacient == nil {
		errMsg := fmt.Sprintf("No pacient found with email: %s", email)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, pacient)
}

func (dController *PacientController) GetPacientByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars["id"]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	log.Printf("[PATIENT] Fetching pacient with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	pacient, err := dController.DbConn.FetchPacientByUserID(ctx, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch pacient with user ID %d: %s", userID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if pacient == nil {
		errMsg := fmt.Sprintf("No pacient found with user ID: %d", userID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, pacient)
}
