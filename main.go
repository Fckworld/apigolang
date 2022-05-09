package main

import (
	"api_mux/api"
	"api_mux/guitar"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("asdsadsa")
	// se usa el \" para sacar las comillas que apareecn el json, creo
}

func main() {
	r := mux.NewRouter()
	a := &api.API{}
	a.RegisterRoutes(r)
	ag := &guitar.GUITAR{}
	ag.RegistrarRutas(r)

	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening...")
	srv.ListenAndServe()
}
