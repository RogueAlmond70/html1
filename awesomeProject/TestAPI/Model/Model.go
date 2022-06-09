package Model

import (
	"TestAPI/View"
	"encoding/json"
	"fmt"
	"net/http"
)

type Record struct {
	Name    string
	Breed   string
	Age     int
	Enquiry string
}

// easy to test
func CreateSampleRecord() Record {
	R001 := Record{
		"Fluffy",
		"Poodle",
		2,
		"Nail clipping",
	}
	return R001
}

func CreateOutputString() string {
	SampleRecord := CreateSampleRecord()
	RecordJSON, err := json.Marshal(SampleRecord)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	OutputString := string(RecordJSON)
	return OutputString
}

// We can test this if we pass it another http.ResponseWriter interface that is just an output to the console
func RetrieveRecord(w http.ResponseWriter, r *http.Request) {
	OutputString := CreateOutputString()
	View.OutPutRecord(w, OutputString)
}
