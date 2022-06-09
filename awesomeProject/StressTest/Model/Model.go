package Model

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func IsNum(s string) error {
	_, err := strconv.Atoi(s)
	return err
}

func IsValid(s string) bool {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	} else {
		if num < 1 {
			return false
		}
	}
	return true
}

func SixtySecondWorkLoad() {
	fmt.Println("Running sixtysecondworkload")
	port := GetPort()
	requests := HowManyRequests()
	MakeRequests(port, requests) //Because the ticker executes actions AFTER the ticker goes off (i.e every 60 seconds, we need to run this MakeRequests() function first before setting the ticker

	ticker := time.NewTicker(60 * time.Second)
	fmt.Println("60 second timer starting...")

	for range ticker.C {
		MakeRequests(port, requests)
		//fmt.Println("60 second timer starting...")
	}
}

func MakeRequests(port string, requests int) {

	fmt.Println("Running MakeRequests")
	for i := 0; i < requests; i++ {

		fmt.Printf("Sending request number %v --  ", i+1)

		r, err := http.Get("http://localhost:" + port + "/retrieve")
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		r.Body.Close() // Closing to prevent memory leak

		output := string(body)
		fmt.Printf("Response %v is: %v \n", i+1, output)
		if err != nil {
			fmt.Println("Error: ", err)
			log.Fatal(err)
		}

	}

}

func HowManyRequests() int {
	fmt.Println("Running HowManyRequests")
	var Requests string
	fmt.Println("How many requests would you like to send per minute? (At least 1)")
	fmt.Scanln(&Requests)
	if err := IsNum(Requests); err != nil {
		fmt.Println("Please enter a valid integer of 1 or greater")
		HowManyRequests()
	} else {
		if IsValid(Requests) == false {
			fmt.Println("Please enter a valid integer between 1 and 1000")
			HowManyRequests()
		}
	}
	var confirmation string
	fmt.Printf("You have specified %v requests. Is this correct? \n", Requests)
	fmt.Scanln(&confirmation)
	if strings.ToLower(confirmation) == "y" || strings.ToLower(confirmation) == "yes" {
		num, _ := strconv.Atoi(Requests)
		fmt.Println("Exiting HowManyRequests")
		return num
	} else {
		HowManyRequests()
	}

	num, _ := strconv.Atoi(Requests)

	return num
}

func GetPort() string {
	fmt.Println("Running GetPort()")
	var PortString string
	var Response string
	fmt.Println("The  default port is 8080. Would you like to use a different port? (Y/N)")
	fmt.Scanln(&Response)
	if strings.ToLower(Response) == "y" || strings.ToLower(Response) == "yes" {
		fmt.Println("Please enter the port number you would like to use: ")
		fmt.Scanln(&PortString)
		_, err := strconv.Atoi(PortString)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			return PortString
		}
	} else {
		PortString = "8080"
	}
	return PortString
}
