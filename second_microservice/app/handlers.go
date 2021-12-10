package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Product struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Sku   string `json:"sku"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!")
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Papitas", "Barcel", "12345"},
		{"Takis", "Barcel", "12346"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(products)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
