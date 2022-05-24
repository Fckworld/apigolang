package main

import (
	"api_mux/api"
	"api_mux/guitar"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"io/ioutil"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ppppppppppp")
	// se usa el \" para sacar las comillas que apareecn el json, creo
}

type task struct {
	ID          int    `json:ID`
	Nombre      string `json:Nombre`
	Descripcion string `json:Descripcion`
}
type allTasks []task

var tasks = allTasks{
	{
		ID:          1,
		Nombre:      "sebastian",
		Descripcion: "ser humano",
	},
}

func crearSer(w http.ResponseWriter, r *http.Request) {
	var nuevoSer task
	infor, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "insertar ser valido")
	}

	json.Unmarshal(infor, &nuevoSer)
	nuevoSer.ID = len(tasks) + 1
	tasks = append(tasks, nuevoSer)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(tasks)

}
func obtenerSeres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //ESTO HACE QUE LAS TAREAS QUE MANDO ESTAN EN FORMATO JSON
	json.NewEncoder(w).Encode(tasks)
}
func obtenerSer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serID, err := strconv.Atoi(vars["id"]) //este id es el que declare en la url

	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}
	for _, task := range tasks {
		if task.ID == serID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}

}

func eliminarSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serID, err := strconv.Atoi(vars["id"]) //este id es el que declare en la url

	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}
	for i, t := range tasks {
		if t.ID == serID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "Ser con id %v fue eliminado", serID)
		}
	}

}

func editarSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serID, err := strconv.Atoi(vars["id"]) //este id es el que declare en la url

	var serditado task //ESTOS SON LOS DATOS QUE QUIERO ACTUALIZAR

	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}

	info, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "insertar info valida")
	}
	json.Unmarshal(info, &serditado)

	for i, t := range tasks {
		if t.ID == serID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			serditado.ID = serID
			tasks = append(tasks, serditado)

			fmt.Fprintf(w, "La tarea con id %v se modifico", serID)
		}
	}
}

func main() {
	router := mux.NewRouter()

	r := mux.NewRouter()
	a := &api.API{}
	a.RegisterRoutes(r)
	ag := &guitar.GUITAR{}
	ag.RegistrarRutas(r)

	r.HandleFunc("/", handleIndex)

	router.HandleFunc("/nuevoser", crearSer).Methods("post")
	router.HandleFunc("/seres", obtenerSeres).Methods("get")
	router.HandleFunc("/ser/{id}", obtenerSer).Methods("get")
	router.HandleFunc("/borra/{id}", eliminarSet).Methods("delete")
	router.HandleFunc("/editar/{id}", editarSet).Methods("put")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Listening...")
	srv.ListenAndServe()
}
