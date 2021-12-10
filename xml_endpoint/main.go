package main

import (
	"encoding/xml"
	"net/http"
)

type Product struct {
	Name  string `xml:"name"`
	Brand string `xml:"brand"`
	Sku   string `xml:"sku"`
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
	w.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(products)
}
