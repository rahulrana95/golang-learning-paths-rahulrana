package main

import "fmt"

func main() {
	fmt.Println("If else if go lang")

	loginCount := 10

	var result = "No"

	if loginCount < 5 {
		result = "Failed"
	} else if loginCount > 5 && loginCount < 10 {
		result = "Failed with 5-10"
	} else {
		result = "Failed  > 10"
	}

	fmt.Println(result)

	// another syntax

	if num := 5; num > 4 {
		fmt.Println("greater then 4")
	} else {
		fmt.Println("less then 4")
	}

}
