package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"first_name"`
	Email string `json:"email"`
}

func MainHAndler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)

}

func main() {
	r := mux.NewRouter()

	// Маршруты
	r.HandleFunc("/products/{id:[0-9]+}", MainHAndler)
	http.Handle("/", r)
	log.Println("Start listening on port :8000")
	http.ListenAndServe(":8000", r)
}
