package main

import (
    "code.google.com/p/go-uuid/uuid"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
)

type Container struct {
    Id uuid.UUID `json:"id"`
}

func main() {
    host, port := configFromEnv()
    listen := fmt.Sprintf("%s:%s", host, port)
    r := mux.NewRouter()
    r.HandleFunc("/", createHandler).Methods("POST")
    log.Fatal(http.ListenAndServe(listen, r))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    id := createUuid()
    e := Container{id}
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

func configFromEnv() (host string, port string) {
    host = os.Getenv("HOST")
    if host == "" {
        host = ""
    }

    port = os.Getenv("PORT")
    if port == "" {
        port = "80"
    }

    return host, port
}