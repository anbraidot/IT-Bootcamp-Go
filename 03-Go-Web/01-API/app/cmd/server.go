package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	// Start the server
	router := chi.NewRouter()

	// Ej. 1
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	// Ej. 2
	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "Hello %s %s", person.FirstName, person.LastName)
	})


	http.ListenAndServe(":8080", router)
}
