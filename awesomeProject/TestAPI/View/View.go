package View

import (
	"fmt"
	"net/http"
)

func OutPutRecord(w http.ResponseWriter, o string) {
	_, err := fmt.Fprintf(w, o)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}
