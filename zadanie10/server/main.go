package main

import (
    "encoding/json"
    "net/http"
    "github.com/rs/cors"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

var products = []Product{
    {ID: 1, Name: "Laptop", Price: 5999.99},
    {ID: 2, Name: "Smartphone", Price: 3999.99},
    {ID: 3, Name: "Smartwatch", Price: 1499.99},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(products)
}

func handlePayment(w http.ResponseWriter, r *http.Request) {
    var paymentData map[string]interface{}
    json.NewDecoder(r.Body).Decode(&paymentData)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Płatność odebrana"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/products", getProducts)
    mux.HandleFunc("/payment", handlePayment)

    handler := cors.Default().Handler(mux)
    http.ListenAndServe(":8080", handler)
}