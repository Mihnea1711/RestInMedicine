package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/middleware"
)

func loadRouter() *mux.Router {
	pacientiHandler := &controllers.Order{}

	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	pacientiRouter := router.PathPrefix("/orders").Subrouter()
	pacientiRouter.HandleFunc("/", pacientiHandler.Create).Methods(http.MethodPost)
	pacientiRouter.HandleFunc("/", pacientiHandler.GetAll).Methods(http.MethodGet)
	pacientiRouter.HandleFunc("/{id:[0-9]+}", pacientiHandler.GetByID).Methods(http.MethodGet)
	pacientiRouter.HandleFunc("/{id:[0-9]+}", pacientiHandler.UpdateByID).Methods(http.MethodPut)
	pacientiRouter.HandleFunc("/{id:[0-9]+}", pacientiHandler.DeleteByID).Methods(http.MethodDelete)

	return router
}
