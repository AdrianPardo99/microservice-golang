package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Init() {
	router := mux.NewRouter()
	/**
	 *   Define routes
	 **/
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	/**
	 * Starting Server
	 **/
	log.Fatal(http.ListenAndServe(":8081", router))
}
