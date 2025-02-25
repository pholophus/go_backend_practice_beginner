package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"
)

type Item struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var items []Item
var nextID = 1

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func basicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || user != "admin" || pass != "password" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	}
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)

	case http.MethodPost:
		var newItem Item

		err := json.NewDecoder(r.Body).Decode(&newItem)

		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		newItem.ID = nextID
		nextID++

		items = append(items, newItem)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func itemHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	//Find item in our slice
	index := -1
	for i, item := range items {
		if item.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items[index])
	case http.MethodPut:
		var updatedItem Item
		err := json.NewDecoder(r.Body).Decode(&updatedItem)
		if err != nil{
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		updatedItem.ID = id
		items[index] = updatedItem
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedItem)
	case http.MethodDelete:
		items = append(items[:index], items[index+1:]...)
        w.WriteHeader(http.StatusNoContent)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/", basicAuth(helloHandler))
	http.HandleFunc("/items", basicAuth(itemsHandler))
	http.HandleFunc("/items/", basicAuth(itemHandler))

	fmt.Println("Server is running on port 8080")

	// http.ListenAndServe(":8080", nil);
	log.Fatal(http.ListenAndServe(":8080", nil))
}