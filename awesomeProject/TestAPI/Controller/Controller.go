package Controller

import (
	"TestAPI/Model"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Homepage Endpoint Hit")
	if err != nil {
		return
	}
}

func MainRouters(r *mux.Router) {
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/retrieve", Model.RetrieveRecord).Methods("GET")
}
