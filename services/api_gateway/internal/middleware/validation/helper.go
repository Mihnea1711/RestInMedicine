package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func logAndRespondWithError(w http.ResponseWriter, statusCode int, logMessage string, err error) {
	log.Printf("[MIDDLEWARE_GATEWAY] %s: %v", logMessage, err)
	response := models.ResponseData{
		Message: logMessage,
		Error:   err.Error(),
	}
	utils.RespondWithJSON(w, statusCode, response)
}

// Function to check the error when decoding an object
func checkErrorOnDecode(err error, w http.ResponseWriter) (bool, int) {
	if err == nil {
		return false, http.StatusOK
	}
	var errMsg string
	var statusCode int

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	switch {
	// Catch any syntax errors in the JSON and send an error message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	case errors.As(err, &syntaxError):
		errMsg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		statusCode = http.StatusUnprocessableEntity

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.Is(err, io.ErrUnexpectedEOF):
		errMsg = "Request body contains badly-formed JSON"
		statusCode = http.StatusUnprocessableEntity

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to a int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		errMsg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		statusCode = http.StatusUnprocessableEntity

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error message and
	// interpolate it in our custom error message. There is an open
	// issue regarding turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		errMsg = fmt.Sprintf("Request body contains unknown field %s", fieldName)
		statusCode = http.StatusUnprocessableEntity

	// An io.EOF error is returned by Decode() if the request body is
	// empty.
	case errors.Is(err, io.EOF):
		errMsg = "Request body must not be empty"
		statusCode = http.StatusBadRequest

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		errMsg = "Request body must not be larger than 1MB"
		statusCode = http.StatusRequestEntityTooLarge

	// Otherwise default to logging the error and sending a 500 Internal
	// Server Error response.
	default:
		errMsg = err.Error()
		statusCode = http.StatusInternalServerError
	}

	log.Printf("[PATIENT_VALIDATION] %s", errMsg)
	return true, statusCode
}

func isContentTypeJSON(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")
	return contentType == "application/json"
}
