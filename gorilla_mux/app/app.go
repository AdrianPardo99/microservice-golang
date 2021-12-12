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
	router.HandleFunc("/hello-world", helloWorld).Methods(http.MethodGet)
	router.HandleFunc("/product", getAllProducts).Methods(http.MethodGet)
	/* Define a parameter number with regular expression */
	router.HandleFunc("/product/{product_id:[0-9]+}", getProduct).Methods(http.MethodGet)
	router.HandleFunc("/product", createProduct).Methods(http.MethodPost)
	/**
	 * Starting Server
	 **/

	log.Fatal(http.ListenAndServe(":8081", router))
}
