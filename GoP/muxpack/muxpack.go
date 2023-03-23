package muxpack

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string
	Id   int
}

func StartMux() {
	mux := mux.NewRouter()

	mux.HandleFunc("/hello", hello)

	// This .Methods is used to explicitly define Get method for an given path!
	mux.HandleFunc("/customers/{customer_id}", getCustomerId).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8000", mux)

}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Returning from the hello from mux function")

}

func getCustomerId(w http.ResponseWriter, r *http.Request) {

	// this mux.Vars returns a map[string]string,
	// which is used to get all the variables keys and the values!
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
