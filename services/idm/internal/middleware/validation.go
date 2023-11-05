package middleware

import (
	"context"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

func ValidateProgramareInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var programare models.Programare

		// dec := json.NewDecoder(r.Body)
		// dec.DisallowUnknownFields()

		// err := dec.Decode(&programare)
		// if checkErrorOnDecode(err, w) {
		// 	return
		// }
		// // Validate the IDProgramare (assuming it should be greater than 0)
		// if programare.IDProgramare < 0 {
		// 	log.Println("[PROGRAMARE] Invalid IDProgramare")
		// 	http.Error(w, "Invalid IDProgramare", http.StatusBadRequest)
		// 	return
		// }

		// // Validate the IDPacient (assuming it should be greater than 0)
		// if programare.IDPacient <= 0 {
		// 	log.Println("[PROGRAMARE] Invalid IDPacient")
		// 	http.Error(w, "Invalid IDPacient", http.StatusBadRequest)
		// 	return
		// }

		// // Validate the IDDoctor (assuming it should be greater than 0)
		// if programare.IDDoctor <= 0 {
		// 	log.Println("[PROGRAMARE] Invalid IDDoctor")
		// 	http.Error(w, "Invalid IDDoctor", http.StatusBadRequest)
		// 	return
		// }

		// // Validate the Date (assuming it should be a valid date)
		// if programare.Date.IsZero() {
		// 	log.Println("[PROGRAMARE] Invalid Date")
		// 	http.Error(w, "Invalid Date", http.StatusBadRequest)
		// 	return
		// }

		// // Validate the Status (assuming it should not be empty)
		// if programare.Status == "" {
		// 	log.Println("[PROGRAMARE] Invalid Status")
		// 	http.Error(w, "Invalid Status", http.StatusBadRequest)
		// 	return
		// }

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_IDM, nil)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func checkErrorOnDecode(err error, w http.ResponseWriter) bool {
// 	if err == nil {
// 		return false
// 	}
// 	var errMsg string

// 	var syntaxError *json.SyntaxError
// 	var unmarshalTypeError *json.UnmarshalTypeError
// 	switch {
// 	// Catch any syntax errors in the JSON and send an error message
// 	// which interpolates the location of the problem to make it
// 	// easier for the client to fix.
// 	case errors.As(err, &syntaxError):
// 		errMsg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

// 	// In some circumstances Decode() may also return an
// 	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
// 	// is an open issue regarding this at
// 	// https://github.com/golang/go/issues/25956.
// 	case errors.Is(err, io.ErrUnexpectedEOF):
// 		errMsg = "Request body contains badly-formed JSON"

// 	// Catch any type errors, like trying to assign a string in the
// 	// JSON request body to a int field in our Person struct. We can
// 	// interpolate the relevant field name and position into the error
// 	// message to make it easier for the client to fix.
// 	case errors.As(err, &unmarshalTypeError):
// 		errMsg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)

// 	// Catch the error caused by extra unexpected fields in the request
// 	// body. We extract the field name from the error message and
// 	// interpolate it in our custom error message. There is an open
// 	// issue regarding turning this into a sentinel error.
// 	case strings.HasPrefix(err.Error(), "json: unknown field "):
// 		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
// 		errMsg = fmt.Sprintf("Request body contains unknown field %s", fieldName)

// 	// An io.EOF error is returned by Decode() if the request body is
// 	// empty.
// 	case errors.Is(err, io.EOF):
// 		errMsg = "Request body must not be empty"

// 	// Catch the error caused by the request body being too large. Again
// 	// there is an open issue regarding turning this into a sentinel
// 	// error at https://github.com/golang/go/issues/30715.
// 	case err.Error() == "http: request body too large":
// 		errMsg = "Request body must not be larger than 1MB"

// 	// Otherwise default to logging the error and sending a 500 Internal
// 	// Server Error response.
// 	default:
// 		errMsg = err.Error()
// 	}

// 	log.Printf("[MIDDLEWARE] %s", errMsg)
// 	http.Error(w, errMsg, http.StatusBadRequest)
// 	return true
// }
