package controllers

import (
	"net/http"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	// Placeholder: Create a new doctor
	w.Write([]byte("Doctor created\n"))
}
