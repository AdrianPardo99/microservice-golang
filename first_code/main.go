package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
	* first form
	* http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	* fmt.Fprint(w, "Hello World!!")
	* })
	**/
	/* Second form */
	http.HandleFunc("/greet", greet)
	http.ListenAndServe("localhost:8081", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}
