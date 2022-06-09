package main

import (
	"TestAPI/Controller"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	fmt.Println("Starting Router")
	Controller.MainRouters(myRouter)

}
