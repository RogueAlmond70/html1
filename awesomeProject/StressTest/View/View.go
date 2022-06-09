package View

import (
	"awesomeProject/StressTest/Model"
	"fmt"
)

func Welcome() {
	fmt.Println("Welcome to our simple Stress tester!")
	Model.SixtySecondWorkLoad()
}
