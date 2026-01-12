package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Khatami. I am an Undergraduate Student")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"image-url"`
}

var productList []Product

// GET /products endpoint to retrieve products
func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreFlight(w, r)

	// if r.Method != "GET" {
	// 	http.Error(w, "Please give me a GET request", 400)
	// 	return
	// }
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	sendData(w, productList, 200)
}

// POST /products endpoint to create a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreFlight(w, r)

	// if r.Method != "POST" {
	// 	http.Error(w, "Please give me a POST request", 400)
	// 	return
	// }
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid product data", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)

	sendData(w, newProduct, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-type, Mahmud")
}

func handlePreFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
	mux.Handle("GET /about", http.HandlerFunc(aboutHandler))
	// mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProducts)))
	mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))

	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))
	mux.Handle("OPTIONS /create-product", http.HandlerFunc(createProduct))

	fmt.Println("Server running on port 3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("There is an error running this server: ", err)
	}
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange contains vitamin C",
		Price:       100,
		ImgUrl:      "redghujicnygygchducbgvusuc",
	}
	prod2 := Product{
		ID:          2,
		Title:       "Mango",
		Description: "Orange contains vitamin C",
		Price:       120,
		ImgUrl:      "redghujicnygygchducbgvusuc",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Orange contains vitamin C",
		Price:       150,
		ImgUrl:      "redghujicnygygchducbgvusuc",
	}

	productList = append(productList, prod1)
	productList = append(productList, prod2)
	productList = append(productList, prod3)
}

// Handle cors through middleware
func corsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-type, Mahmud")

		next.ServeHTTP(w, r)
	}
	handler := http.HandlerFunc(handleCors)
	return handler
}
