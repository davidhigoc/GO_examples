package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const get string = "GET"
const post string = "POST"
const put string = "PUT"
const delete string = "DELETE"

// Interface que comunica los metodos que e asocian a una estructura
// ? se necesita un constructor
type InterfaceGeneral interface {
	HandlerGene(w http.ResponseWriter, r *http.Request)
	HandlerOne(w http.ResponseWriter, r *http.Request)
}

type HandlerGeneralStruct struct {
}

// CONSTRUCTOR DE INTERFACE
func BuildInterfaceGeneral() InterfaceGeneral {
	return HandlerGeneralStruct{}
}

func (h HandlerGeneralStruct) HandlerGene(w http.ResponseWriter, r *http.Request) {
	// Distribución de metodos
	switch r.Method {
	case get:
		fmt.Println("petición Get-ON")
		m := make(map[string][]string)
		fmt.Println(m)
		slice := []string{"Pedro", "Benito", "Dolores", "Cantasia", "Sexto"}
		m["Users"] = append(m["Users"], slice...)
		json.NewEncoder(w).Encode(m)
	case post:
		m := make(map[string]interface{})
		slice := []string{"Pedro", "Benito", "Dolores", "Cantasia", "Sexto"}
		m["Users"] = slice
		m["Estations"] = []string{"Verano", "Invierno", "Otoño", "Primavera"}
		m["Version"] = 1
		fmt.Println("petición Post-ON")
		json.NewEncoder(w).Encode(m)
	case put:
		fmt.Println("petición Put-ON")
		json.NewEncoder(w).Encode(put)
	case delete:
		fmt.Println("petición Delete-ON")
		json.NewEncoder(w).Encode(delete)

	}
}

func (h HandlerGeneralStruct) HandlerOne(w http.ResponseWriter, r *http.Request) {
	// Distribución de metodos
	switch r.Method {
	case get:
		// Identificamos parametro URL
		id, encontrado := mux.Vars(r)["id"]
		// Evaluamos parametro URL
		if encontrado {
			resp := "Se recibio el id = " + id
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(nil)
		}
		json.NewEncoder(w).Encode(get)
	case post:
		fmt.Println("petición Post-ON")
		json.NewEncoder(w).Encode(post)
	case put:
		fmt.Println("petición Put-ON")
		json.NewEncoder(w).Encode(put)
	case delete:
		fmt.Println("petición Delete-ON")
		json.NewEncoder(w).Encode(delete)

	}
}
