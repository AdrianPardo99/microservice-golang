package app

import (
	"log"
	"net/http"
)

func Init() {
	mux := http.NewServeMux()
	/**
	 *   Define routes
	 **/
	mux.HandleFunc("/hello-world", helloWorld)
	mux.HandleFunc("/get-products", getAllProducts)

	/**
	 * Starting Server
	 **/

	log.Fatal(http.ListenAndServe("0.0.0.0:8081", mux))
}
