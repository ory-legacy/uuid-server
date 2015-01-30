package main

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Id struct {
	Id uuid.UUID `json:"string"`
}

func main() {
	portPtr := flag.String("port", "80", "The port to listen on")
	hostPtr := flag.String("host", "", "The adress to listen on. Leave empty to listen on all interfaces")
	flag.Parse()
	listen := fmt.Sprintf("%s:%s", *hostPtr, *portPtr)
	r := mux.NewRouter()
	r.HandleFunc("/", createHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(listen, r))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	id := createUuid()
	e := Id{id}
	j, err := json.Marshal(e)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func createUuid() uuid.UUID {
	return uuid.NewRandom()
}
