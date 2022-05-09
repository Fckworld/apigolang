package guitar

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (g *GUITAR) RegistrarRutas(r *mux.Router) {
	r.HandleFunc("/guitarras", g.getGuitars).Methods(http.MethodGet)
	r.HandleFunc("/guitarras", g.addGuitar).Methods(http.MethodPost)
	r.HandleFunc("/guitarras/{id}", g.getGuitar).Methods(http.MethodGet)
	r.HandleFunc("/dguitarra/{id}", g.deleteGuitar).Methods(http.MethodDelete)
	r.HandleFunc("/uguitarra/{id}", g.updateGuitar).Methods(http.MethodPut)
}
