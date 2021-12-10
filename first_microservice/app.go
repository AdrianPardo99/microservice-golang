package main

import (
	"log"
	"net/http"
)

func Init() {
	/**
	 *   Define routes
	 **/
	http.HandleFunc("/hello-world", helloWorld)
	http.HandleFunc("/get-products", getAllProducts)

	/**
	 * Starting Server
	 **/

	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
