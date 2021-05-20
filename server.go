package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/kellemnegasi/restapi-with-mux/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "up and running...")
	})

	router.HandleFunc("/posts", routes.GetPost).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")
	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
