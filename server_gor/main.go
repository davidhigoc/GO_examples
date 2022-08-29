package main

import (
	"fmt"
	"go_pru1/server_gor/handlers"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Iniciando Servidor")

	handlergeneral := handlers.BuildInterfaceGeneral()

	// Instancia para nuevo Router
	router := mux.NewRouter()

	// Handlers
	router.HandleFunc("/", handlergeneral.HandlerGene)
	router.HandleFunc("/{id}", handlergeneral.HandlerOne)

	server := http.Server{
		Addr:              ":3000", // adress - Abrir el servicio por el puerto
		Handler:           router,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	fmt.Println("Escuchando...")
	server.ListenAndServe()

}
