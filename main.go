package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id   int
	Name string
}

func customers(w http.ResponseWriter, r *http.Request) {
	c := Customer{Id: 1, Name: "Elias"}
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/customers", customers).Methods("Get").Name("customers")

	err := http.ListenAndServe("go-project-crt-643934-dev.apps.sandbox-m2.ll9k.p1.openshiftapps.com/", router)
	if err != nil {
		log.Fatal(err)
	}
}
