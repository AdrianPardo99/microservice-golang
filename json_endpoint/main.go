package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Sku   string `json:"sku"`
}

func main() {
	http.HandleFunc("/get-products", getAllProducts)
	http.ListenAndServe("localhost:8081", nil)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Papitas", "Barcel", "12345"},
		{"Takis", "Barcel", "12346"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
