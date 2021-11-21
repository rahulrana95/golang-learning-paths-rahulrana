package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case  go lang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Its 1")
	case 2:
		fmt.Println("Its 2")
		fallthrough // next will execute
	case 3:
		fmt.Println("Its 3")
	case 4:
		fmt.Println("Its 4")
	case 5:
		fmt.Println("Its 5")
	case 6:
		fmt.Println("Its 6")
	}
}
