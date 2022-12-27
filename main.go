package main

import (
	"Caruio/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.DOMCar(r)
	r.HandleFunc("/", dome)
	log.Fatal(http.ListenAndServe("localhost:2023", r))
}
func dome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dome")
}
