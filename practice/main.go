package main

import (
	"fmt"
)

func main() {

	//test := []int{0, 2, 3, 4, 5}
	//fmt.Println(find_min(test))
	test := []int{7, 3, 7, 3, 1, 3, 4, 1}
	//fmt.Println(Solution(7))
	fmt.Println(Solution(test))

}

func find_min(A []int) int {
	var result int = 0
	for i := 1; i < len(A); i++ {
		if result > A[i] {
			result = A[i]
		}
	}
	return result
}

// Given a non-empty array consisting of N integers, returns the length of the shortest vacation that allows
// you to visit lall the offered locations
func Solution(A []int) int {
	tripLengths := []int{}
	//test := []int{7, 3, 7, 3, 1, 3, 4, 1}
	var output *int
	for i := 0; i < len(A); i++ {
		fmt.Println("Starting Outer loop")
		locations := []int{}
		var days = 0

		for j := i; j < len(A); j++ {
			days++
			if len(locations) == 0 {
				locations = append(locations, j)
				fmt.Println("Appending")
			}
			fmt.Println("Starting first loop")
			for k := 0; k < len(locations); k++ {
				fmt.Println("starting inner loop")
				if j == k {
					break
					fmt.Println("Breakingut of firrst loop")
				}
				locations = append(locations, j)
				if len(locations) == 4 {
					break
					fmt.Println("Breaking out of second loop")
				}
			}
		}
		tripLengths = append(tripLengths, days)

	}
	// Return the smallest value in the slice
	var min = 9999
	for w := 0; w < len(tripLengths); w++ {
		if w == 0 || w < min {
			min = w
			fmt.Println(w)
		}
		output = &min

	}

	return *output
}

/*
// Must generate a slice of "random" integers greater than 0
func Solution(N int) []int {
	rand.Seed(time.Now().UnixNano())
	var output []int
	for i := 0; i < N; i++ {
		ourInt := rand.Intn(1000-1+1) + 1
		output = append(output, ourInt)
	}
	return output
}
*/
