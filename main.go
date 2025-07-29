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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Please give me a GET request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

// POST /products endpoint to create a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, "Please give me a POST request", 400)
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
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-product", createProduct)

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
