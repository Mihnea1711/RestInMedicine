package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Retrieve all consultatii
func (cController *ConsultatieController) GetConsultatii(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve consultatii.")

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch all consultatii from the database
	consultatii, err := cController.DbConn.FetchAllConsultatii(ctx, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultatii: %s\n", errMsg)
		response := models.ResponseData{
			Error: errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[CONSULTATION] Successfully fetched all consultatii")

	// Serialize the consultatii to JSON and send the response
	response := models.ResponseData{
		Payload: consultatii,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve a consultatie by ID
func (cController *ConsultatieController) GetConsultatieByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve a consultatie by ID.")

	vars := mux.Vars(r)
	id := vars[utils.FETCH_CONSULTATIE_BY_ID_PARAMETER]

	// Convert the ID to a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response := models.ResponseData{
			Error: "Invalid consultatie ID",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch the consultatie by ID from the database
	consultatie, err := cController.DbConn.FetchConsultatieByID(ctx, objectID)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultatie by ID: %s\n", errMsg)
		response := models.ResponseData{
			Error: errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the consultatie exists
	if consultatie == nil {
		response := models.ResponseData{
			Error: "Consultatie not found",
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	log.Printf("[CONSULTATION] Successfully fetched consultatie %s", consultatie.IDConsultatie)
	// Serialize the consultatie to JSON and send the response
	response := models.ResponseData{
		Payload: consultatie,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve consultatii by doctor ID
func (cController *ConsultatieController) GetConsultatiiByDoctorID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve consultatii by Doctor ID.")

	vars := mux.Vars(r)
	doctorID, err := strconv.Atoi(vars[utils.FETCH_CONSULTATIE_BY_DOCTOR_ID_PARAMETER])
	if err != nil {
		response := models.ResponseData{
			Error: "Invalid Doctor ID",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	log.Printf("READ: %d", doctorID)

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch consultatii by Doctor ID from the database
	consultatii, err := cController.DbConn.FetchConsultatiiByDoctorID(ctx, doctorID, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultatii by Doctor ID: %s\n", errMsg)
		response := models.ResponseData{
			Error: errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[CONSULTATION] Successfully fetched consultatii of doctor %d", doctorID)
	// Serialize the consultatii to JSON and send the response
	response := models.ResponseData{
		Payload: consultatii,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve consultatii by pacient ID
func (cController *ConsultatieController) GetConsultatiiByPacientID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve consultatii by Pacient ID.")

	vars := mux.Vars(r)
	pacientID, err := strconv.Atoi(vars[utils.FETCH_CONSULTATIE_BY_PACIENT_ID_PARAMETER])
	if err != nil {
		response := models.ResponseData{
			Error: "Invalid Pacient ID",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch consultatii by Pacient ID from the database
	consultatii, err := cController.DbConn.FetchConsultatiiByPacientID(ctx, pacientID, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultatii by Pacient ID: %s\n", errMsg)
		response := models.ResponseData{
			Error: errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[CONSULTATION] Successfully fetched consultatii of pacient %d", pacientID)
	// Serialize the consultatii to JSON and send the response
	response := models.ResponseData{
		Payload: consultatii,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve consultatii by date
func (cController *ConsultatieController) GetConsultatiiByDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve consultatii by date.")

	// Use mux.Vars to get the date parameter from the route
	vars := mux.Vars(r)
	dateStr := vars[utils.FETCH_CONSULTATIE_BY_DATE_PARAMETER]

	// Parse the date string into a time.Time object
	date, err := time.Parse(utils.TIME_FORMAT, dateStr)
	if err != nil {
		errResponse := models.ResponseData{
			Error: "Invalid date format",
		}
		log.Printf("[CONSULTATION] Failed to convert date string: %s\n", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, errResponse)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch consultatii by date from the database
	consultatii, err := cController.DbConn.FetchConsultatiiByDate(ctx, date, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		errResponse := models.ResponseData{
			Error: errMsg,
		}
		log.Printf("[CONSULTATION] Failed to fetch consultatii by date: %s\n", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, errResponse)
		return
	}

	log.Printf("[CONSULTATION] Successfully fetched consultatii from %s", date)
	// Serialize the consultatii to JSON and send the response
	response := models.ResponseData{
		Payload: consultatii,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (cController *ConsultatieController) GetFilteredConsultatii(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve filtered consultatii.")

	// Extract query filter params
	filter := utils.ExtractQueryParams(r)

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch filtered consultatii from the database
	consultatii, err := cController.DbConn.FetchConsultatiiByFilter(ctx, filter, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch filtered consultatii: %s\n", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, errMsg)
		return
	}

	if len(consultatii) != 0 {
		log.Printf("[CONSULTATION] Successfully fetched filtered consultatii")
	} else {
		log.Printf("[CONSULTATION] No consultatii found with the filter: %v", filter)
	}
	// Serialize the filtered consultatii to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, consultatii)
}
