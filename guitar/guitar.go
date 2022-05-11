package guitar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GUITAR struct {
}
type guitarra struct {
	ID      int    `json:ID gorm:"primaryKey"`
	Name    string `json:Name`
	Content string `json:Content`
}

type allGuitars []guitarra

var guitars = allGuitars{
	{
		ID:      1,
		Name:    "AAAAA",
		Content: "ZZZZZ",
	},
	{
		ID:      2,
		Name:    "BBBBB",
		Content: "YYYYY",
	},
	{
		ID:      3,
		Name:    "CCCCC",
		Content: "XXXXX",
	},
	{
		ID:      4,
		Name:    "DDDDD",
		Content: "WWWWW",
	},
}

func (g *GUITAR) getGuitars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guitars)
}

func (g *GUITAR) addGuitar(w http.ResponseWriter, r *http.Request) {

	var newGuitar guitarra
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &newGuitar)
	newGuitar.ID = len(guitars) + 1
	guitars = append(guitars, newGuitar)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGuitar)
}

func (g *GUITAR) getGuitar(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	idParam := pathParams["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	index := id - 1
	if index < 0 || index > len(guitars)-1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(guitars[index])
}

func (g *GUITAR) deleteGuitar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guitarID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "id invalida")
		return
	}

	for i, t := range guitars {
		if t.ID == guitarID {
			guitars = append(guitars[:i], guitars[i+1:]...)
			fmt.Fprintf(w, "La guitarra %v se elimino", guitarID)
		}
	}
}
func (g *GUITAR) updateGuitar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guitarID, err := strconv.Atoi(vars["id"])
	var updatedTask guitarra

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range guitars {
		if t.ID == guitarID {
			guitars = append(guitars[:i], guitars[i+1:]...)

			updatedTask.ID = t.ID
			guitars = append(guitars, updatedTask)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", guitarID)
		}
	}

}
