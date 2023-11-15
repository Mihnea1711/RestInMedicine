package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) GetDoctors(w http.ResponseWriter, r *http.Request) {
	log.Println("[DOCTOR] Fetching all doctors...")

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctors, err := dController.DbConn.FetchDoctors(ctx)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, doctors)
}

func (dController *DoctorController) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", idStr)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch doctor with ID %d: %s", id, err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with ID: %d", id)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, doctor)
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
		errMsg := fmt.Sprintf("Failed to fetch doctor with email %s: %s", email, err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with email: %s", email)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, doctor)
}

func (dController *DoctorController) GetDoctorByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars["id"]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByID(ctx, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch doctor with user ID %d: %s", userID, err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with user ID: %d", userID)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, doctor)
}
