// internal/controllers/delete.go

package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Placeholder: Delete doctor with given ID
	w.Write([]byte("Deleted doctor with ID: " + id))
}
